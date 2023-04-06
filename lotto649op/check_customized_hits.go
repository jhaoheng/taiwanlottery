package lotto649op

import (
	"fmt"
)

/*
- 檢查輸入的號碼, 在歷史中, 中獎機率
- hit_num_count, 設定至少要幾個號碼中
*/

func (op *Lotto649OP) CheckCustomizedHits(hit_num_count int, nums ...string) (hit_lottery []Lotto649OPData) {
	if len(nums) < 6 {
		fmt.Println("選號數字過少, 最少要六個")
		return
	}
	//
	choice := map[string]bool{}
	for _, num := range nums {
		choice[num] = true
	}

	hit_lottery = []Lotto649OPData{}
	for _, data := range op.Datas {
		hits := 0
		if _, ok := choice[data.Num_1]; ok {
			hits++
		}
		if _, ok := choice[data.Num_2]; ok {
			hits++
		}
		if _, ok := choice[data.Num_3]; ok {
			hits++
		}
		if _, ok := choice[data.Num_4]; ok {
			hits++
		}
		if _, ok := choice[data.Num_5]; ok {
			hits++
		}
		if _, ok := choice[data.Num_6]; ok {
			hits++
		}
		if _, ok := choice[data.NumSpecial]; ok {
			hits++
		}
		//
		if hits >= hit_num_count {
			hit_lottery = append(hit_lottery, data)
		}
	}
	//
	return hit_lottery
}
