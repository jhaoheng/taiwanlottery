package module

import (
	"fmt"
	"strings"
	"time"

	"github.com/jhaoheng/taiwanlottery/model"
)

/*
- 多線程將資料寫入 db
*/

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
