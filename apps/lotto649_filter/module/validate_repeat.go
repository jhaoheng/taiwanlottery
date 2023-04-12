package module

import (
	"fmt"
	"strconv"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) CheckRepeat(datas [][]int) {
	map_data := map[string]struct{}{}
	for _, data := range datas {
		if len(data) != 6 {
			panic("沒有 6 個數字")
		}
		num_1 := strconv.Itoa(data[0])
		num_2 := strconv.Itoa(data[1])
		num_3 := strconv.Itoa(data[2])
		num_4 := strconv.Itoa(data[3])
		num_5 := strconv.Itoa(data[4])
		num_6 := strconv.Itoa(data[5])
		//
		text := fmt.Sprintf("%v,%v,%v,%v,%v,%v", num_1, num_2, num_3, num_4, num_5, num_6)

		if _, ok := map_data[text]; ok {
			fmt.Printf("重複數據: %v\n", text)
			break
		} else {
			map_data[text] = struct{}{}
		}
	}
}
