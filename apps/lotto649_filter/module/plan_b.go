package module

/*
- 排除掉連續 N 個數字, ex: 若 N=4 則 (1,2,3,4,9,11) 會被過濾掉
*/

type IPlanB interface {
	GetCombinations() (combinations [][]int)
}

type PlanB struct {
}

func NewPlanB() IPlanB {
	return &PlanB{}
}

func (plan *PlanB) GetCombinations() (combinations [][]int) {
	combinations = [][]int{}
	numbers := []int{1, 2, 3, 4, 5, 6, 7} // 需要排列的數字
	size := 6                             // 每個排列的大小
	consecutive := 5                      // 要求的連續數字數量

	//
	tmp_combinations := plan.generateCombinations(numbers, size)
	for _, c := range tmp_combinations {
		if plan.hasConsecutive(c, consecutive) {
			combinations = append(combinations, c)
		}
	}
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
