package plan

import (
	"fmt"
	"time"
)

/*
經過 plan_c 的計算
在 plan_b 的做法, 最好取連續 3 個數字, 的排列組合
*/

/*
- 產生指定的連續數字組合

ex:
數字 [1,2,3,4,5,6,7] 的組合有以下可能
[1 2 3 4 5 6]
[1 2 3 4 5 7]
[1 3 4 5 6 7]
[2 3 4 5 6 7]
*/

type IPlanB interface {
	// [產生指定連續數字的組合]
	GetCombinations(min_num, max_num, consecutive, size int) (combinations [][]int)
}

type PlanB struct {
}

func NewPlanB() IPlanB {
	return &PlanB{}
}

/*
[產生指定連續數字的組合]
consecutive: 要求的連續數字數量
size: 每個排列的大小
*/
func (plan *PlanB) GetCombinations(min_num, max_num, consecutive, size int) (combinations [][]int) {
	fmt.Println("=== PlanB ===")
	start := time.Now()
	combinations = [][]int{}
	numbers := []int{} // 需要排列的數字
	for i := min_num; i <= max_num; i++ {
		numbers = append(numbers, i)
	}

	//
	tmp_combinations := plan.generateCombinations(numbers, size)
	for _, c := range tmp_combinations {
		if plan.hasConsecutive(c, consecutive) {
			combinations = append(combinations, c)
		}
	}
	fmt.Printf("總共有: %v, 執行時間: %v\n", len(combinations), -time.Until(start))
	return
}

// 產生指定大小的排列組合
func (plan *PlanB) generateCombinations(numbers []int, size int) [][]int {
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

// 檢查排列中是否存在指定數量的連續數字
func (plan *PlanB) hasConsecutive(combination []int, consecutive int) bool {
	count := 1
	for i := 0; i < len(combination)-1; i++ {
		if combination[i]+1 == combination[i+1] {
			count++
		} else {
			count = 1
		}

		if count == consecutive {
			return true
		}
	}

	return false
}
