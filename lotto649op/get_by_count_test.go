package lotto649op

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_GetByCount(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	start := time.Now().In(loc).AddDate(-3, 0, 0)
	end := time.Now().In(loc).AddDate(0, 0, -2)
	result, start_data, end_data := NewLotto649OP(raw_results).GetNumCount(start, end)

	fmt.Println("起始資料=>", func() string {
		b, _ := json.MarshalIndent(start_data, "", "	")
		return string(b)
	}())
	fmt.Println("最後資料=>", func() string {
		b, _ := json.MarshalIndent(end_data, "", "	")
		return string(b)
	}())

	result.OrderNumCount()

	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))
}
