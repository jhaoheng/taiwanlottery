package lotto649op

import (
	"fmt"
)

/*
- 排除 => 歷史 頭獎 與 二獎 號碼 => 中五個號碼以上
*/
func (op *Lotto649OP) Excluded_1(sets PossibleSets) (result PossibleSets) {
	type empty struct{}
	var member empty

	var new_set PossibleSets = [][]int{}

	// 創建一個 map，用來快速檢查 set 是否符合 op.Datas 的條件
	setMaps := make([]map[string]empty, len(sets))
	for i, set := range sets {
		setMaps[i] = make(map[string]empty)
		for _, num := range set {
			setMaps[i][fmt.Sprintf("%02d", num)] = member
		}
	}

	for i, set := range sets {
		save_flag := true
		for _, data := range op.Datas {
			hit := 0
			if _, ok := setMaps[i][data.Num_1]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_2]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_3]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_4]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_5]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_6]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.NumSpecial]; ok {
				hit++
			}
			if hit >= 5 {
				save_flag = false
				break
			}
		}
		if save_flag {
			new_set = append(new_set, set)
		}
	}
	return new_set
}

/*
- 排除 => 歷史頭獎 取 5 個號碼, 再搭配 1 個號碼
- sets 數字, ascending
- remove_sets, ascending
*/
func (op *Lotto649OP) Excluded_2(sets PossibleSets, remove_sets []string) (result PossibleSets) {

	// for _, set := range sets {
	// 	sort.Ints(set)
	// }

	// var new_set PossibleSets = [][]int{}
	// for _, set := range sets {
	// 	save_flag := true
	// 	for _, remove_set := range remove_sets {
	// 		hit := 0
	// 		if _, ok := map_set[remove_set[0]]; ok {
	// 			hit++
	// 		}
	// 		if _, ok := map_set[remove_set[0]]; ok {
	// 			hit++
	// 		}
	// 		if _, ok := map_set[remove_set[0]]; ok {
	// 			hit++
	// 		}
	// 		if _, ok := map_set[remove_set[0]]; ok {
	// 			hit++
	// 		}
	// 		if _, ok := map_set[remove_set[0]]; ok {
	// 			hit++
	// 		}
	// 		if _, ok := map_set[remove_set[0]]; ok {
	// 			hit++
	// 		}
	// 		if hit == 6 {
	// 			save_flag = false
	// 		}
	// 	}
	// }
	return
}

/*
- 排除 => 上一次頭獎的 7 個號碼
*/
func (op *Lotto649OP) Excluded_3(sets PossibleSets) (result PossibleSets) {
	return
}
