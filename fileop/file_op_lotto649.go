package fileop

import (
	"strings"
)

type Lotto649CSV struct {
	GameName       string
	SerialID       string // 期別
	Date           string // 開獎日期
	TotalSales     string // 銷售總額
	TotalAmount    string // 銷售注數
	TheRewardMoney string // 總獎金
	Num_1          string
	Num_2          string
	Num_3          string
	Num_4          string
	Num_5          string
	Num_6          string
	Num_special    string
}

/*
- split_space: 切割欄位的符號, ex: `,` or `	`
*/
func (fop *FileOP) ParsedLotto649CSV(split_space string) (results []Lotto649CSV, err error) {
	// // 取得結構所需 keys
	// struct_keys, err := fop.get_struct_keys(Lotto649CSV{})
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("all keys => %v\n\n", struct_keys)
	//
	results = []Lotto649CSV{}
	for index, eachline := range fop.lines {
		if index == 0 {
			continue
		}
		values := strings.Split(eachline, split_space)
		results = append(results, Lotto649CSV{
			GameName:       values[0],
			SerialID:       values[1],
			Date:           values[2],
			TotalSales:     values[3],
			TotalAmount:    values[4],
			TheRewardMoney: values[5],
			Num_1:          values[6],
			Num_2:          values[7],
			Num_3:          values[8],
			Num_4:          values[9],
			Num_5:          values[10],
			Num_6:          values[11],
			Num_special:    values[12],
		})
	}
	return
}
