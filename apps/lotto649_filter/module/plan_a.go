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
消耗時間: 1.169109939s
總共=> 891610
*/

type IPlanA interface {
	// 取得 7 取 5 排列組合
	GetCombinations(all_hits []lotto649op.Lotto649OPData) (combinations [][]int)
	// 補齊 六個號碼
	FillTo6(combinations [][]int) (results [][]int, results_map map[string]struct{})
	//
	RunFilter(guess_sets map[string]struct{}, filter_combinations [][]int) map[string]struct{}
}

type PlanA struct {
	Start time.Time
}

func NewPlanA() IPlanA {
	return &PlanA{
		Start: time.Now(),
	}
}

/*
- 總共有: 21756, 執行時間: 19.678175ms
*/
func (plan *PlanA) GetCombinations(all_hits []lotto649op.Lotto649OPData) (combinations [][]int) {
	fmt.Println("=== PlanA.GetCombinations() ===")
	combinations = [][]int{}
	for _, hit := range all_hits {
		/*
			1. 組合出適當的數組
			2. 從資料庫中判斷是否有相同號碼
			3. 若有找到該號碼, 則刪除
		*/
		// fmt.Printf("=== 目前處理(%v): %v ===\n", index, hit.SerialID)

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
	fmt.Printf("總共有: %v, 執行時間: %v\n", len(combinations), -time.Until(plan.Start))
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

/*
- 補齊六組號碼
*/
func (plan *PlanA) FillTo6(combinations [][]int) (results [][]int, results_map map[string]struct{}) {
	fmt.Println("=== PlanA.FillTo6() ===")
	//
	results, results_map = FillTo6(combinations)
	//
	fmt.Printf("消耗時間: %v\n", -time.Until(plan.Start))
	return results, results_map
}

/*
-
*/
func (plan *PlanA) RunFilter(guess_sets map[string]struct{}, filter_combinations [][]int) map[string]struct{} {
	for _, data := range filter_combinations {
		num_1 := strconv.Itoa(data[0])
		num_2 := strconv.Itoa(data[1])
		num_3 := strconv.Itoa(data[2])
		num_4 := strconv.Itoa(data[3])
		num_5 := strconv.Itoa(data[4])
		num_6 := strconv.Itoa(data[5])
		text := fmt.Sprintf("%v,%v,%v,%v,%v,%v", num_1, num_2, num_3, num_4, num_5, num_6)

		//
		delete(guess_sets, text)
	}
	return guess_sets
}
