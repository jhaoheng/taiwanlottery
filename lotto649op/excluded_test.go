package lotto649op

import (
	"fmt"
	"testing"
)

func Test_Excluded(t *testing.T) {
	/*
		測試 112000039 期的數字預測
	*/

	results := NewLotto649OP(raw_results).GetAllSets()
	fmt.Println(len(results))
}
