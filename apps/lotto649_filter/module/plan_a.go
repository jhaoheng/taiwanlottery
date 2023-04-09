package module

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
)

/*
- 取得 中獎號碼, 7 取 5 排列組合
- all_hits 可設定時間區間, 取得資料
*/

type IPlanA interface {
	FillTo6(combinations [][]int) [][]int
	GetCombinations(all_hits []lotto649op.Lotto649OPData) (combinations [][]int)
}

type PlanA struct {
}

func NewPlanA() IPlanA {
	return &PlanA{}
}

// 補齊六組號碼
func (plan *PlanA) FillTo6(combinations [][]int) [][]int {
	//
	add_num := func(nums []int) [][]int {
		new_nums := [][]int{}
		num_map := map[int]struct{}{}
		for _, num := range nums {
			num_map[num] = struct{}{}
		}
		//
		for i := 1; i <= 49; i++ {
			if _, ok := num_map[i]; !ok {
				tmp := append(nums, i)
				new_nums = append(new_nums, tmp)
			}
		}
		return new_nums
	}

	//
	result := [][]int{}
	for _, nums := range combinations {
		result = append(result, add_num(nums)...)
	}
	return result
}

// -
func (plan *PlanA) GetCombinations(all_hits []lotto649op.Lotto649OPData) (combinations [][]int) {
	fmt.Println("=== PlanA ===")
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
		sort.Ints(numbers)
		combinations = append(combinations, plan.M_Get_N(numbers, 5)...)
	}
	fmt.Printf("總共有: %v, 執行時間: %v\n", len(combinations), -time.Until(start))
	return
}

/*
- M 個數字中, 取 N 個數字, 進行排列組合
- ex: numbers=[]int{1,2,3,4,5,6,7}, count=5 => 7 取 5 進行排列組合
*/
func (plan *PlanA) M_Get_N(numbers []int, count int) (combinations [][]int) {
	combinations = [][]int{} // 最後組合
	plan.generateCombinations(&combinations, []int{}, numbers, count)

	// // 輸出排列組合
	// fmt.Println("排列組合：")
	// for _, combination := range combinations {
	// 	fmt.Println(combination)
	// }
	return combinations
}

/**/
func (plan *PlanA) generateCombinations(combinations *[][]int, current []int, remaining []int, count int) {
	if count == 0 {
		*combinations = append(*combinations, current)
		return
	}
	for i := 0; i < len(remaining); i++ {
		newCurrent := append(current, remaining[i])
		newRemaining := append([]int{}, remaining[i+1:]...)
		plan.generateCombinations(combinations, newCurrent, newRemaining, count-1)
	}
}
