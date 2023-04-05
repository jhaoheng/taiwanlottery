package lotto649op

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_GetNextDataByTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	the_time := time.Now().In(loc).AddDate(0, 0, -2)
	data := NewLotto649OP(raw_results).GetNextDataByTime(the_time)

	fmt.Println("查詢時間 =>", the_time)
	b, _ := json.MarshalIndent(data, "", "	")
	fmt.Println(string(b))
}
