package module

/*
- M 個數字中, 取 N 個數字, 進行排列組合
- ex: numbers=[]int{1,2,3,4,5,6,7}, count=5 => 7 取 5 進行排列組合
*/
func M_Get_N(numbers []int, count int) (combinations [][]int) {
	combinations = [][]int{} // 最後組合
	generateCombinations(&combinations, []int{}, numbers, count)

	// // 輸出排列組合
	// fmt.Println("排列組合：")
	// for _, combination := range combinations {
	// 	fmt.Println(combination)
	// }
	return combinations
}

/**/
func generateCombinations(combinations *[][]int, current []int, remaining []int, count int) {
	if count == 0 {
		*combinations = append(*combinations, current)
		return
	}
	for i := 0; i < len(remaining); i++ {
		newCurrent := append(current, remaining[i])
		newRemaining := append([]int{}, remaining[i+1:]...)
		generateCombinations(combinations, newCurrent, newRemaining, count-1)
	}
}
