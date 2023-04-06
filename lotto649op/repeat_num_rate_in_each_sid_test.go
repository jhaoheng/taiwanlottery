package lotto649op

import (
	"testing"
	"time"
)

func Test_GetRepeatNumRateInEachSID(t *testing.T) {

	loc, _ := time.LoadLocation("Asia/Taipei")
	repeat := 3
	start, _ := time.ParseInLocation("2006-01-02", "2014-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-01-01", loc)
	NewLotto649OP(raw_results).GetRepeatNumRateInEachSID(repeat, start, end)
}
