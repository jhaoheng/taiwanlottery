package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/module"
	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

var raw_results []model.Lottery

var op lotto649op.ILotto649OP

func main() {
	// model.IsDebug = true
	err := model.ConnMySQL()
	if err != nil {
		panic(err)
	}

	raw_results, _ = model.NewLottery().FindAll()
	op = lotto649op.NewLotto649OP(raw_results)
	// module.GetAllPossiblility(op) // 取得所有組數, 並寫入資料庫

	/*
		[目前策略] 指定時間區間
		1. 取得指定區間的中獎資料
		2. 取得推估的數據(消耗時間: 21.100764774s)
		3. 取得 plan_d 的組合 (消耗時間: 3.832917736s), 進行過濾
		4. 取得 plan_a 的數字（消耗時間: 1.160604563s），進行過濾
		5. 取得 plan_e 的數字（消耗時間: 53.774µs），進行過濾
		6. 取得可能的數字


		=== 開始 GetAllSets ===
		消耗時間: 21.100764774s
		=== 開始 PlanD.GetConsecutiveSets() ===
		消耗時間: 22.598µs
		=== PlanD.FillTo6() ===
		消耗時間: 4.89455652s
		=== PlanA.GetCombinations() ===
		總共有: 21735, 執行時間: 7.085448ms
		=== PlanA.FillTo6() ===
		消耗時間: 976.342363ms
		=== PlanE.GetSpecificNums() ===
		消耗時間: 44.477µs
		最後剩下 => 2899068
	*/

	// 1. 取得指定區間的中獎資料
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, _ := time.ParseInLocation("2006-01-02", "1973-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-04", loc)
	hits := []lotto649op.Lotto649OPData{}
	for _, hit := range op.GetLotto649OPDatas() {
		if hit.Date.Unix() < start.Unix() || hit.Date.Unix() > end.Unix() {
			continue
		}
		hits = append(hits, hit)
	}
	fmt.Printf("全部 his 有: %v\n", len(hits))

	// 2. 取得推估的數據(消耗時間: 21.100764774s)
	_, all_sets_map := module.NewAllSets().Get()
	fmt.Println("原始組合 =>", len(all_sets_map))

	// 3. 取得 plan_d 的組合 (消耗時間: 3.832917736s), 進行過濾
	all_sets_map = func() map[string]struct{} {
		plan_d := module.NewPlanD()
		filter_combinations := plan_d.GetConsecutiveSets(3)
		filter_combinations, _ = plan_d.FillTo6(filter_combinations)
		all_sets_map = plan_d.RunFilter(all_sets_map, filter_combinations)
		fmt.Println("剩下 =>", len(all_sets_map))
		return all_sets_map
	}()

	// 4. 取得 plan_a 的數字（消耗時間: 222.021544ms），進行過濾
	all_sets_map = func() map[string]struct{} {
		plan_a := module.NewPlanA()
		filter_combinations := plan_a.GetCombinations(hits)
		filter_combinations, _ = plan_a.FillTo6(filter_combinations)
		all_sets_map = plan_a.RunFilter(all_sets_map, filter_combinations)
		fmt.Println("剩下 =>", len(all_sets_map))
		return all_sets_map
	}()

	/* 5. 取得 plan_e 的數字（消耗時間: 53.774µs），進行過濾
	- 只取得最後一次中獎數字
	*/
	all_sets_map = func() map[string]struct{} {
		loc, _ := time.LoadLocation("Asia/Taipei")
		start, _ := time.ParseInLocation("2006-01-02", "2023-04-04", loc)
		end, _ := time.ParseInLocation("2006-01-02", "2023-04-04", loc)
		plan_e := module.NewPlanE()
		filter_combinations := plan_e.GetSpecificNums(hits, start, end)
		fmt.Println("準備過濾的數字: ", filter_combinations)
		all_sets_map = plan_e.RunFilter(all_sets_map, filter_combinations)
		fmt.Println("剩下 =>", len(all_sets_map))
		return all_sets_map
	}()

	// 6. 取得可能的數字，寫入資料庫
	WriteToDB(all_sets_map)

}

func WriteToDB(all_sets_map map[string]struct{}) {
	batch_write_size := 2000
	table_name := "filter_3"
	filtered := model.NewLotto649Filtered(table_name)
	fmt.Printf("=== 開始寫入 database, table_name: %v ===\n", table_name)
	start := time.Now()

	// 整理批次
	fmt.Println("批次整理")
	save_group_keys := [][]string{}
	save_keys := []string{}
	for key := range all_sets_map {
		tmps := strings.Split(key, ",")
		if len(tmps) != 6 {
			err := fmt.Errorf("字數錯誤: %v", tmps)
			panic(err)
		}
		if len(save_keys) != batch_write_size {
			save_keys = append(save_keys, key)
		} else {
			save_group_keys = append(save_group_keys, save_keys)
			save_keys = []string{}
		}
	}
	if len(save_keys) != 0 {
		save_group_keys = append(save_group_keys, save_keys) // 最後一組
	}

	fmt.Printf("總共 %v 組\n", len(save_group_keys))

	/*
		準備發送
	*/
	task_chan := make(chan []string, 200)
	done_chan := make(chan struct{}, 200)

	doSomething := func() {
		for {
			texts := <-task_chan
			// fmt.Printf("收到任務: %v\n", len(texts))
			datas := []model.Lotto649Filtered{}
			for _, text := range texts {
				datas = append(datas, model.Lotto649Filtered{
					Nums: text,
				})
			}
			filtered.CreateInBatch(datas, batch_write_size)
			done_chan <- struct{}{}
		}
	}

	for i := 0; i < 10; i++ {
		go doSomething()
	}

	//
	go func() {
		for _, save_keys := range save_group_keys {
			task_chan <- save_keys
		}
	}()

	//
	finish_group := 0
	for range done_chan {
		finish_group++

		if finish_group%1000 == 0 {
			fmt.Printf("====> 完成: %v, 消耗時間: %v \n", finish_group, -time.Until(start))
		}

		if finish_group == len(save_group_keys) {
			close(done_chan)
			close(task_chan)
			break
		}
	}

	//
	fmt.Printf("消耗時間: %v\n", -time.Until(start))
}
