package module

import (
	"fmt"
	"sort"
	"strconv"
)

func FillTo6(combinations [][]int) (results [][]int, results_map map[string]struct{}) {
	//
	add_num := func(nums []int) [][]int {
		//
		num_map := map[int]struct{}{}
		for _, num := range nums {
			num_map[num] = struct{}{}
		}
		//
		new_nums := [][]int{}
		for i := 1; i <= 49; i++ {
			if _, ok := num_map[i]; !ok {
				tmp := []int{}
				tmp = append(tmp, nums...)
				tmp = append(tmp, i)
				sort.Ints(tmp)
				// fmt.Println(tmp)
				new_nums = append(new_nums, tmp)
			}
		}
		return new_nums
	}

	//
	tmp_results := [][]int{}
	for {
		if len(combinations[0]) == 6 {
			break
		}
		tmp_results = [][]int{}
		for _, nums := range combinations {
			// fmt.Println(nums)
			tmp := add_num(nums)
			// fmt.Println(tmp)
			tmp_results = append(tmp_results, tmp...)
			// break
		}
		combinations = tmp_results
	}

	//
	results = [][]int{}
	result_map := map[string]struct{}{}
	for _, data := range tmp_results {
		num_1 := strconv.Itoa(data[0])
		num_2 := strconv.Itoa(data[1])
		num_3 := strconv.Itoa(data[2])
		num_4 := strconv.Itoa(data[3])
		num_5 := strconv.Itoa(data[4])
		num_6 := strconv.Itoa(data[5])
		//
		text := fmt.Sprintf("%v,%v,%v,%v,%v,%v", num_1, num_2, num_3, num_4, num_5, num_6)
		if _, ok := result_map[text]; !ok {
			result_map[text] = struct{}{}
			results = append(results, data)
		}
	}
	return results, result_map
}
