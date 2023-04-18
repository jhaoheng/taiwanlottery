package lotto649op

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_CheckCustomizedHits_1(t *testing.T) {
	texts := []string{
		"04, 07, 08, 25, 34, 45",
		"04, 08, 21, 24, 31, 39",
		"04, 15, 25, 33, 39, 44",
		"04, 17, 28, 30, 40, 46",
		"04, 19, 33, 38, 45, 48",
		"05, 20, 29, 30, 36, 37",
		"06, 08, 34, 36, 41, 48",
		"06, 15, 36, 38, 39, 41",
		"06, 17, 37, 40, 43, 44",
		"06, 20, 24, 31, 32, 39",
		"07, 08, 15, 34, 39, 41",
		"07, 17, 24, 29, 31, 48",
		"07, 17, 29, 30, 34, 36",
		"07, 18, 20, 37, 39, 40",
		"07, 21, 36, 38, 40, 41",
		"08, 15, 18, 19, 21, 48",
		"08, 15, 28, 32, 37, 39",
		"08, 28, 32, 38, 41, 46",
		"15, 18, 19, 24, 37, 39",
		"15, 19, 20, 32, 33, 48",
		"17, 19, 21, 29, 31, 32",
		"17, 21, 25, 28, 29, 38",
		"17, 24, 28, 32, 39, 41",
		"18, 21, 32, 34, 43, 46",
		"20, 30, 36, 43, 44, 46",
		"11, 25, 30, 31, 37, 45",
		"25, 28, 32, 34, 40, 41",
		"04, 11, 20, 25, 31, 33", // 新增
		"11, 13, 21, 37, 41, 42", // 新增
	}

	datas := [][]string{}
	for _, text := range texts {
		items := strings.Split(text, ", ")
		datas = append(datas, items)
	}

	// datas = append(datas, []string{"06", "07", "10", "16", "34", "41", "09"}) // 8 組, 這是我買的
	// datas = append(datas, []string{"03", "13", "17", "18", "31", "49"}) // 8 組, 這是我買的
	// datas = append(datas, []string{"02", "07", "09", "11", "33", "36"}) // 8 組
	// datas = append(datas, []string{"02", "08", "11", "14", "34", "36"}) // 10 組
	// // datas = append(datas, []string{"01", "03", "05", "08", "11", "13", "14", "20", "21", "30", "33", "36"})
	// datas = append(datas, []string{"03", "13", "16", "19", "21", "23", "01"})
	// datas = append(datas, []string{"09", "16", "22", "33", "41", "34"})
	// datas = append(datas, []string{"04", "11", "20", "25", "32", "39", "30"})

	for _, data := range datas {
		result := NewLotto649OP(raw_results).CheckCustomizedHits(5, data...)
		b, _ := json.MarshalIndent(result, "", "	")
		fmt.Println(string(b))
		fmt.Printf("%v, 歷史上，總共中獎有 => %v 組\n", data, len(result))
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

/*
- 測試指定的（前幾期, ex: 1 個月）中獎號碼，出現在下一期的機率
*/
func Test_CheckCustomizedHitsWithTime(t *testing.T) {

	loc, _ := time.LoadLocation("Asia/Taipei")
	op := NewLotto649OP(raw_results)

	//
	next := 2
	// 取得下一期的中獎號碼
	fmt.Printf("取得下 %v 期的中獎號碼\n", next)
	nums := func() []string {
		tmp := map[string]struct{}{}
		current_time, _ := time.ParseInLocation("2006-01-02", "2023-04-01", loc)
		datas := op.GetNextDataByTime(current_time, 1)
		fmt.Println(func() string {
			b, _ := json.MarshalIndent(datas, "", "	")
			return string(b)
		}())
		//
		for _, data := range datas {
			tmp[data.Num_1] = struct{}{}
			tmp[data.Num_2] = struct{}{}
			tmp[data.Num_3] = struct{}{}
			tmp[data.Num_4] = struct{}{}
			tmp[data.Num_5] = struct{}{}
			tmp[data.Num_6] = struct{}{}
			tmp[data.NumSpecial] = struct{}{}
		}
		//
		results := []string{}
		for num := range tmp {
			results = append(results, num)
		}
		return results
	}()

	//
	start, _ := time.ParseInLocation("2006-01-02", "2023-03-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-01", loc)
	fmt.Printf("跟前期比較, %v - %v\n", start.Format("2006-01-02"), end.Format("2006-01-02"))
	hits := op.CheckCustomizedHitsWithTime(start, end, 4, nums...)
	fmt.Println("對中幾組 :", len(hits))
	for _, hit := range hits {
		fmt.Println(func() string {
			b, _ := json.MarshalIndent(hit, "", "	")
			return string(b)
		}())
	}
}
