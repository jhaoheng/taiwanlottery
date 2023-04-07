package lotto649op

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_CheckCustomizedHits_1(t *testing.T) {

	datas := [][]string{}
	// datas = append(datas, []string{"06", "07", "10", "16", "34", "41", "09"}) // 8 組, 這是我買的
	// datas = append(datas, []string{"03", "13", "17", "18", "31", "49"}) // 8 組, 這是我買的
	// datas = append(datas, []string{"02", "07", "09", "11", "33", "36"}) // 8 組
	// datas = append(datas, []string{"02", "08", "11", "14", "34", "36"}) // 10 組
	// // datas = append(datas, []string{"01", "03", "05", "08", "11", "13", "14", "20", "21", "30", "33", "36"})
	// datas = append(datas, []string{"03", "13", "16", "19", "21", "23", "01"})
	// datas = append(datas, []string{"09", "16", "22", "33", "41", "34"})
	datas = append(datas, []string{"03", "04", "16", "29", "32", "39"})

	for _, data := range datas {
		result := NewLotto649OP(raw_results).CheckCustomizedHits(6, data...)

		b, _ := json.MarshalIndent(result, "", "	")
		fmt.Println(string(b))
		fmt.Printf("歷史上，總共中獎有 => %v 組\n", len(result))
	}
}

/*
- 指定使用歷史中獎號碼，查看中獎率多高
*/
func Test_CheckCustomizedHits_With_HitLotto(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, _ := time.ParseInLocation("2006-01-02", "2014-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-05", loc)
	op := NewLotto649OP(raw_results)
	hit_nums := op.GetLotto649OPDatasWithFactor(start, end)
	fmt.Println("驗證數量=>", len(hit_nums))

	//
	for _, hit_num := range hit_nums {
		result := NewLotto649OP(raw_results).CheckCustomizedHits(6, hit_num...)

		// b, _ := json.MarshalIndent(result, "", "	")
		// fmt.Println(string(b))
		if len(result) != 0 {
			fmt.Printf("[%v]，總共中獎有 => %v 組\n", hit_num, len(result))
		}
	}
}
