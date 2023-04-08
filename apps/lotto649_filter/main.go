package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

var raw_results []model.Lottery

var op lotto649op.ILotto649OP

func main() {
	err := model.ConnSQLite("../../sql.db")
	if err != nil {
		panic(err)
	}

	raw_results, _ = model.NewLottery().FindAll()
	op = lotto649op.NewLotto649OP(raw_results)

	//
	GetAllPossiblility()
	//
	// Filter_1(sets)
}

/*
- 取得所有組數
- 寫入資料庫
- 總共需要時間: 2m6.642164592s
*/
func GetAllPossiblility() {
	fmt.Println("=== 開始取得所有組數 ===")
	//
	start := time.Now()
	//
	results := op.GetAllSets()
	fmt.Printf("所有可能組數: %v, 處理時間: %v\n", len(results), -time.Until(start))

	//
	fmt.Println("寫入資料格式化...")
	writes := []model.Lotto649AllSets{}
	for index, result := range results {
		if index%1000000 == 0 {
			fmt.Println("目前已格式化 =>", index)
		}
		writes = append(writes, model.Lotto649AllSets{
			Nums: func() string {
				if len(result) != 6 {
					panic("nums length not 6")
				}
				sort.Slice(result, func(i, j int) bool {
					return result[i] < result[j]
				})
				return fmt.Sprintf("%02d,%02d,%02d,%02d,%02d,%02d", result[0], result[1], result[2], result[3], result[4], result[5])
			}(),
		})
	}

	//
	fmt.Printf("格式化結束, 組數: %v, 處理時間: %v\n", len(writes), -time.Until(start))
	fmt.Printf("開始寫入資料庫\n")
	if err := model.NewLotto649AllSets().DeleteAll(); err != nil {
		panic(err)
	}
	if err := model.NewLotto649AllSets().CreateInBatch(writes, 100); err != nil {
		panic(err)
	}
	fmt.Printf("done!!, 處理時間: %v\n\n", -time.Until(start))
}

/*
- 會花比較久的時間, 第一次執行約一小時
- 只能排除掉約 一百萬筆
- 總共需要時間: 55 分鐘
*/
func Filter_1(sets lotto649op.PossibleSets) {
	//
	// sets = sets[:len(sets)/10000]
	//
	start := time.Now()
	//
	results := op.Excluded_1(sets)
	fmt.Println("剩下組數 =>", len(results))

	//
	fmt.Println("整理寫入資料格式化..., 總共筆數 =>", len(results), ", ", -time.Until(start))
	writes := []model.Lotto649AllSets{}
	for index, result := range results {
		if index%1000 == 0 {
			fmt.Println("目前 =>", index)
		}
		writes = append(writes, model.Lotto649AllSets{
			Nums: func() string {
				if len(result) != 6 {
					panic("nums length not 6")
				}
				sort.Slice(result, func(i, j int) bool {
					return result[i] < result[j]
				})
				return fmt.Sprintf("%02d,%02d,%02d,%02d,%02d,%02d", result[0], result[1], result[2], result[3], result[4], result[5])
			}(),
		})
	}

	//
	fmt.Println("開始寫入資料庫, 總共筆數 =>", len(writes), ", ", -time.Until(start))
	if err := model.NewLotto649AllSets().DeleteAll(); err != nil {
		panic(err)
	}
	model.NewLotto649AllSets().CreateInBatch(writes, 1000)
	fmt.Println("done!!")
	fmt.Println("filter_1 =>", -time.Until(start))
}
