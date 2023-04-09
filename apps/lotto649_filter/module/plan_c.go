package module

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
)

/*
- 經過測試
- 連續兩個號碼的機率有 669/1037
- 連續三個號碼的機率有 91/1037
- 連續四個號碼的機率有 8/1037
*/

/*
- 取得已中獎的號碼, 當中有連續數字
- 取得後，過濾掉該數字
*/

type IPlanC interface {
	GetconsecutiveHit(all_hits []lotto649op.Lotto649OPData, consecutive int) (combinations [][]int, exceptions []lotto649op.Lotto649OPData)
}

type PlanC struct {
}

func NewPlanC() IPlanC {
	return &PlanC{}
}

func (plan *PlanC) GetconsecutiveHit(all_hits []lotto649op.Lotto649OPData, consecutive int) (combinations [][]int, exceptions []lotto649op.Lotto649OPData) {
	fmt.Println("=== PlanC ===")
	start := time.Now()
	tmp_combinations := [][]int{}
	for index, hit := range all_hits {
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
		// 判斷是否存在指定的連續數字
		if !plan.hasConsecutiveNumbers(numbers, consecutive) {
			tmp_combinations = append(tmp_combinations, numbers)
		} else {
			exceptions = append(exceptions, hit)
		}
	}

	// 存活下來的數字，7 取 6，排列組合
	for _, tmp := range tmp_combinations {
		combinations = plan.generateCombinations(tmp, 6)
	}

	fmt.Printf("總共有: %v, 執行時間: %v\n", len(combinations), -time.Until(start))
	return
}

// 判斷是否有連續數字
func (plan *PlanC) hasConsecutiveNumbers(group []int, m int) bool {
	count := 1
	for i := 1; i < len(group); i++ {
		if group[i] == group[i-1]+1 {
			count++
		} else {
			count = 1
		}
		if count == m {
			return true
		}
	}
	return false
}

// 產生指定大小的排列組合
func (plan *PlanC) generateCombinations(numbers []int, size int) [][]int {
	if size == 1 {
		result := [][]int{}
		for _, n := range numbers {
			result = append(result, []int{n})
		}
		return result
	}

	result := [][]int{}
	for i, n := range numbers {
		for _, sub := range plan.generateCombinations(numbers[i+1:], size-1) {
			result = append(result, append([]int{n}, sub...))
		}
	}
	return result
}
