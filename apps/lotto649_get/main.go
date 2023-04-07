package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

var raw_results []model.Lottery

var op lotto649op.ILotto649OP

func main() {
	err := model.ConnSQLite("../../sql.db")
	if err != nil {
		panic(err)
	}

	raw_results, _ = model.NewLottery().FindAll()

	//
	op = lotto649op.NewLotto649OP(raw_results)
	sets := op.GetAllSets()
	fmt.Println("組數 =>", len(sets))

	//
	Filter_1(sets)
}

/*
- 會花比較久的時間, 第一次執行約一小時
- 只能排除掉約 一百萬筆
*/
func Filter_1(sets lotto649op.PossibleSets) {
	//
	sets = sets[:len(sets)/1000]
	//
	start := time.Now()
	results := op.Excluded_1(sets)
	fmt.Println("剩下組數 =>", len(results))
	fmt.Println("filter_1 =>", time.Until(start))

	//
	BreakLineTag := "\r\n"
	var csv string
	for _, result := range results {
		csv = csv + fmt.Sprintf("%02d", result[0]) + ","
		csv = csv + fmt.Sprintf("%02d", result[1]) + ","
		csv = csv + fmt.Sprintf("%02d", result[2]) + ","
		csv = csv + fmt.Sprintf("%02d", result[3]) + ","
		csv = csv + fmt.Sprintf("%02d", result[4]) + ","
		csv = csv + fmt.Sprintf("%02d", result[5]) + BreakLineTag
	}

	//
	filepath := "./filter_1.csv"
	if _, err := os.Stat(filepath); err != nil {
		fmt.Println("file not exist")
	} else {
		if err := os.Remove(filepath); err != nil {
			panic(err)
		}
	}
	if err := os.WriteFile(filepath, []byte(csv), 0777); err != nil {
		panic(err)
	}
}
