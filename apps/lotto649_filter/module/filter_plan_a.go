package module

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

/*
- 需要先將原始資料複製到資料表(lotto649_filtered)中
- 過濾掉中獎號碼的數組 (對中 >=5)
*/

type IFilterPlanA interface {
	StartFilter(all_hits []lotto649op.Lotto649OPData)
}

type FilterPlanA struct {
}

/*
 */
func New_FilterPlanA() IFilterPlanA {
	return &FilterPlanA{}
}

func (filter *FilterPlanA) StartFilter(all_hits []lotto649op.Lotto649OPData) {
	start := time.Now()

	del_count := 0
	for index, hit := range all_hits {
		/*
			1. 組合出適當的數組
			2. 從資料庫中判斷是否有相同號碼
			3. 若有找到該號碼, 則刪除
		*/
		fmt.Printf("=== 目前處理(%v): %v ===\n", index, hit.SerialID)

		// 產生排列組合
		n1, _ := strconv.Atoi(hit.Num_1)
		n2, _ := strconv.Atoi(hit.Num_2)
		n3, _ := strconv.Atoi(hit.Num_3)
		n4, _ := strconv.Atoi(hit.Num_4)
		n5, _ := strconv.Atoi(hit.Num_5)
		n6, _ := strconv.Atoi(hit.Num_6)
		n7, _ := strconv.Atoi(hit.NumSpecial)
		numbers := []int{
			n1, n2, n3, n4, n5, n6, n7,
		}
		combinations := [][]int{}
		filter.generateCombinations(&combinations, []int{}, numbers, 5)
		// // 輸出排列組合
		// fmt.Println("排列組合：")
		// for _, combination := range combinations {
		// 	fmt.Println(combination)
		// }

		//
		for _, numbers := range combinations {
			fmt.Printf("=> 查詢號碼: %v, ", numbers)
			//
			text := ""
			for _, n := range numbers {
				text = text + "%" + fmt.Sprintf("%02d", n) + "%"
			}
			finds, err := model.NewLotto649Filtered().FindNumsLike([]string{text})
			if err != nil {
				panic(err)
			}
			fmt.Printf("找到 %v\n", len(finds))
			if len(finds) != 0 {
				if err := model.NewLotto649Filtered().BatchDelete(finds); err != nil {
					panic(err)
				}
				del_count = del_count + len(finds)
			}
		}
	}
	fmt.Printf("總共刪除: %v, 執行時間: %v\n", del_count, -time.Until(start))
}

/**/
func (filter *FilterPlanA) generateCombinations(combinations *[][]int, current []int, remaining []int, count int) {
	if count == 0 {
		*combinations = append(*combinations, current)
		return
	}
	for i := 0; i < len(remaining); i++ {
		newCurrent := append(current, remaining[i])
		newRemaining := append([]int{}, remaining[i+1:]...)
		filter.generateCombinations(combinations, newCurrent, newRemaining, count-1)
	}
}
