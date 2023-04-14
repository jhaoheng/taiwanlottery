package module

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

/*
- 取得所有可能的組合 6 個號碼
*/

type IAllSets interface {
	Get() (all_sets [][]int, all_sets_map map[string]struct{})
}

type AllSets struct {
}

func NewAllSets() IAllSets {
	return &AllSets{}
}

func (sets *AllSets) Get() (all_sets [][]int, all_sets_map map[string]struct{}) {
	fmt.Println("=== 開始 GetAllSets ===")
	start := time.Now()
	max_num := 49
	num_count := 6
	all_sets = [][]int{}
	//
	sets.generateSequence(map[int]bool{}, max_num, num_count, []int{}, &all_sets)
	//
	all_sets_map = map[string]struct{}{}
	for _, set := range all_sets {
		sort.Ints(set)
		num_1 := strconv.Itoa(set[0])
		num_2 := strconv.Itoa(set[1])
		num_3 := strconv.Itoa(set[2])
		num_4 := strconv.Itoa(set[3])
		num_5 := strconv.Itoa(set[4])
		num_6 := strconv.Itoa(set[5])

		text := fmt.Sprintf("%v,%v,%v,%v,%v,%v", num_1, num_2, num_3, num_4, num_5, num_6)
		if _, ok := all_sets_map[text]; !ok {
			all_sets_map[text] = struct{}{}
		}
	}
	//
	fmt.Printf("消耗時間: %v\n", -time.Until(start))
	return all_sets, all_sets_map
}

func (sets *AllSets) generateSequence(excluded_nums map[int]bool, max_num int, num_count int, nums []int, results *[][]int) {
	// 如果生成了指定數量的數字，則將其添加到結果中
	if len(nums) == num_count {
		*results = append(*results, append([]int{}, nums...))
		return
	}

	// 生成下一個數字
	for i := 1; i <= max_num; i++ {
		if !excluded_nums[i] && (len(nums) == 0 || i > nums[len(nums)-1]) {
			excluded_nums[i] = true
			nums = append(nums, i)
			sets.generateSequence(excluded_nums, max_num, num_count, nums, results)
			nums = nums[:len(nums)-1]
			excluded_nums[i] = false
		}
	}
}
