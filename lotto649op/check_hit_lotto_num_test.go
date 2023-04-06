package lotto649op

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_CheckHitLotto(t *testing.T) {
	results := NewLotto649OP(raw_results).CheckHitLotto(5)

	//
	for _, result := range results {
		b, _ := json.MarshalIndent(result, "", "	")
		fmt.Println(string(b))
	}

	fmt.Println("總共重複次數 =>", len(results))
}
