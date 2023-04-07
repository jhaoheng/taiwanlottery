package lotto649op

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_AccumulatedDatas(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, _ := time.ParseInLocation("2006-01-02", "2014-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-03-01", loc)
	accumulated_datas := NewLotto649OP(raw_results).AccumulatedDatasByTime(start, end, 1)

	b, _ := json.MarshalIndent(accumulated_datas, "", "	")
	fmt.Println(string(b))
}

func Test_AccumulatedDatas_ExportCSV_1(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, _ := time.ParseInLocation("2006-01-02", "2014-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-07", loc)
	accumulated_datas := NewLotto649OP(raw_results).AccumulatedDatasByTime(start, end, 0)

	// b, _ := json.MarshalIndent(accumulated_datas, "", "	")
	// fmt.Println(string(b))

	filename, csv := accumulated_datas.ExportCSV_1()
	// fmt.Println(filename)
	// fmt.Println(csv)

	filepath := "./accumulated_output/" + filename
	if _, err := os.Stat(filepath); err != nil {
		fmt.Println("file not exist")
	} else {
		if err := os.Remove(filepath); err != nil {
			panic(err)
		}
	}
	//
	if err := os.WriteFile(filepath, []byte(csv), os.FileMode(0777)); err != nil {
		panic(err)
	}
}
