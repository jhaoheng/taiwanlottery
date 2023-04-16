package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
- 確定移除的數字越準確，推估到的排列越正確
*/

var remove_arrs = []int{}    // 確定移除的數字
var remove_special = []int{} // 根據情境, 要移除的特別數字
var n_sets = 28              // 取得幾組數字

func checkConsecutive(data []int) bool {
	for i := 0; i < len(data)-2; i++ {
		if data[i]+1 == data[i+1] && data[i+1]+1 == data[i+2] {
			return true
		}
	}
	return false
}

func combinationsUtil(arr []int, data []int, start int, end int, index int, results *[][]int) {
	if index == 6 {
		if !checkConsecutive(data) {
			result := []int{}
			for i := 0; i < 6; i++ {
				result = append(result, data[i])
			}
			*results = append(*results, result)
		}
		return
	}

	for i := start; i <= end && end-i+1 >= 6-index; i++ {
		data[index] = arr[i]
		combinationsUtil(arr, data, i+1, end, index+1, results)
	}
}

func combinations(arr []int) [][]int {
	results := [][]int{}

	n := len(arr)
	data := make([]int, 6)
	combinationsUtil(arr, data, 0, n-1, 0, &results)
	return results
}

func main() {
	inferential_arrs := get_inference_arr()
	fmt.Println("移除數字後, 預估的數字...")
	fmt.Printf("%+v\n", inferential_arrs)
	fmt.Println()
	//
	datas := combinations(inferential_arrs)
	fmt.Println("總共可能出現組數", len(datas))

	//
	fmt.Printf("隨機出現的 %v 組排列\n", n_sets)
	rand_output(datas)
}

func rand_output(datas [][]int) {
	rand.Seed(time.Now().UnixNano())

	// 隨機取出 28 個數字
	selectedNums := [][]int{}
	for i := 0; i < n_sets; i++ {
		randIndex := rand.Intn(len(datas))
		selectedNums = append(selectedNums, datas[randIndex])
		datas = append(datas[:randIndex], datas[randIndex+1:]...)
	}

	sort.Slice(selectedNums, func(i, j int) bool {
		return selectedNums[i][0] < selectedNums[j][0]
	})

	for _, v := range selectedNums {
		fmt.Println(v)
	}
}

func get_inference_arr() []int {
	remove_map := map[int]struct{}{}
	for _, v := range remove_arrs {
		remove_map[v] = struct{}{}
	}
	for _, v := range remove_special {
		remove_map[v] = struct{}{}
	}

	//
	datas := []int{}
	for i := 1; i <= 49; i++ {
		if _, ok := remove_map[i]; !ok {
			datas = append(datas, i)
		}
	}
	return datas
}
