package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
- 從最後過濾的數字組中，取得所有個別數字
- 從所有個別數字中，隨機給予 N 個號碼
- 該號碼所有組合，必須要是過濾後的數字組之一
*/

var arr = []int{4, 5, 6, 7, 8, 15, 17, 18, 19, 20, 21, 24, 25, 28, 29, 30, 31, 32, 33, 34, 36, 37, 38, 39, 40, 41, 43, 44, 45, 46, 48}
var n_sets = 28 // 取得幾組數字

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
	datas := combinations(arr)

	fmt.Println(len(datas))
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
