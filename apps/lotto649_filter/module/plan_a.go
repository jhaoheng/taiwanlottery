package module

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
)

/*
- 取得 中獎號碼, 7 取 5 排列組合
- all_hits 可設定時間區間, 取得資料
*/

type IPlanA interface {
	GetCombinations(all_hits []lotto649op.Lotto649OPData) (combinations [][]int)
}

type PlanA struct {
}

func NewPlanA() IPlanA {
	return &PlanA{}
}

func (plan *PlanA) GetCombinations(all_hits []lotto649op.Lotto649OPData) (combinations [][]int) {
	start := time.Now()
	combinations = [][]int{}
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
		combinations = append(combinations, M_Get_N(numbers, 5)...)
	}
	fmt.Printf("執行時間: %v\n", -time.Until(start))
	return
}
