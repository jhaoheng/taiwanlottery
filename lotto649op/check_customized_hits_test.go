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
	datas = append(datas, []string{"01", "11", "18", "20", "21", "35"})

	for _, data := range datas {
		result := NewLotto649OP(raw_results).CheckCustomizedHits(5, data...)

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
	start, _ := time.ParseInLocation("2006-01-02", "2022-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-05", loc)
	op := NewLotto649OP(raw_results)
	hit_nums := op.GetLotto649OPDatasAndReplaceOne(start, end)
	fmt.Println("驗證數量=>", len(hit_nums))

	//
	for index, hit_num := range hit_nums {
		results := NewLotto649OP(raw_results).CheckCustomizedHits(6, hit_num...)

		// b, _ := json.MarshalIndent(result, "", "	")
		// fmt.Println(string(b))
		if len(results) != 0 {
			fmt.Printf("=== %v:%v ===\n", index, hit_num)
			for _, result := range results {
				fmt.Printf("=> %v, %v\n", result.Date.Format("2006-01-02"), result.SerialID)
			}
			fmt.Printf("總共中獎有 => %v 組\n", len(results))
			fmt.Println()
		}
	}
}
