package lotto649op

import (
	"fmt"
	"testing"
)

func Test_CheckHits(t *testing.T) {

	datas := [][]string{}
	datas = append(datas, []string{"06", "07", "10", "16", "34", "41"}) // 8 組
	datas = append(datas, []string{"02", "07", "09", "11", "33", "36"}) // 8 組
	datas = append(datas, []string{"02", "08", "11", "14", "34", "36"}) // 10 組
	datas = append(datas, []string{"01", "03", "05", "08", "11", "13", "14", "20", "21", "30", "33", "36"})

	for _, data := range datas {
		result := NewLotto649OP(raw_results).CheckHits(6, data...)

		// b, _ := json.MarshalIndent(result, "", "	")
		// fmt.Println(string(b))
		fmt.Printf("歷史上，總共中獎有 => %v 組\n", len(result))
	}
}
