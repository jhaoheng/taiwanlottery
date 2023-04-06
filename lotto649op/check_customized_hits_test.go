package lotto649op

import (
	"fmt"
	"testing"
)

func Test_CheckCustomizedHits(t *testing.T) {

	datas := [][]string{}
	datas = append(datas, []string{"06", "07", "10", "16", "34", "41"}) // 8 組, 這是我買的
	datas = append(datas, []string{"03", "13", "17", "18", "31", "49"}) // 8 組, 這是我買的
	datas = append(datas, []string{"02", "07", "09", "11", "33", "36"}) // 8 組
	datas = append(datas, []string{"02", "08", "11", "14", "34", "36"}) // 10 組
	// datas = append(datas, []string{"01", "03", "05", "08", "11", "13", "14", "20", "21", "30", "33", "36"})
	datas = append(datas, []string{"03", "13", "16", "19", "21", "23", "01"})
	datas = append(datas, []string{"11", "18", "20", "21", "35", "37", "08"})

	for _, data := range datas {
		result := NewLotto649OP(raw_results).CheckCustomizedHits(4, data...)

		// b, _ := json.MarshalIndent(result, "", "	")
		// fmt.Println(string(b))
		fmt.Printf("歷史上，總共中獎有 => %v 組\n", len(result))
	}
}
