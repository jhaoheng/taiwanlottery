package module

import (
	"fmt"
	"sort"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

/*
- 取得所有組數
- 寫入資料庫
- 處理時間: 12m53.183714852s
*/
func GetAllPossiblility(op lotto649op.ILotto649OP) {
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
	if err := model.NewLotto649AllSets().CreateInBatch(writes, 1000); err != nil {
		panic(err)
	}
	fmt.Printf("done!!, 處理時間: %v\n\n", -time.Until(start))
}
