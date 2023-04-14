package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/flowactions"
	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/module"
	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/plan"
	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
	"gorm.io/gorm"
)

/*
	raw_results, _ = model.NewLottery().FindAll()
	op = lotto649op.NewLotto649OP(raw_results)
	module.GetAllPossiblility(op) // 取得所有組數, 並寫入資料庫
*/

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

	//
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, err := time.ParseInLocation("2006-01-02", "1973-01-01", loc)
	if err != nil {
		panic(err)
	}
	// end, err := time.ParseInLocation("2006-01-02", "2023-02-17", loc)
	// if err != nil {
	// 	panic(err)
	// }

	_, all_sets_map := module.NewAllSets().Get()
	final_result := [][]plan.PlanFCountRankOnlyHitIndex{}
	start_id := 768
	for i := start_id; i < len(raw_results); i++ {
		data, err := model.NewLottery().SetID(int64(i)).Take()
		if err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		end := data.Date
		flow_a := flowactions.NewFlowA(op, start, end).SetAllSetsMap(all_sets_map)
		results := flow_a.Run().GetRankAndExportOnlyHitIndexes()
		final_result = append(final_result, results)
	}

	/*
		建立 csv
	*/
	fmt.Println("=== 開始建立 csv ===")
	BreakLineTag := "\r\n"

	csv := ""

	for i := 1; i <= 49; i++ {
		csv = csv + "," + fmt.Sprintf("%v", i)
	}
	csv = csv + BreakLineTag

	for i := 0; i < len(final_result); i++ {
		csv = csv + final_result[i][0].HitSerialID
		for j := 0; j < len(final_result[i]); j++ {
			value := ""
			if final_result[i][j].Hit {
				value = "1"
			}
			csv = csv + "," + fmt.Sprintf("%v", value)
		}
		csv = csv + BreakLineTag
	}

	// fmt.Println(csv)
	os.WriteFile("test.csv", []byte(csv), 0777)

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
