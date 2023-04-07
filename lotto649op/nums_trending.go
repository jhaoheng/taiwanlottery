package lotto649op

import (
	"fmt"
	"time"
)

/*
- 所有個別號碼，在時間中的出現趨勢
*/

type NumsTrendingResults []NumsTrendingResult

type NumsTrendingResult struct {
	Date time.Time // 開獎日期, ex: 2006/01/01 15:04:05
	Num  string
}

/*
- filename: nums_trending_{最後一筆資料的時間}
*/
func (op *Lotto649OP) ExportNumsTrending() (filename, csv string) {
	BreakLineTag := "\r\n"

	//
	keys := []string{}
	csv = func() string {
		// 建立 title
		title := "Date"
		for i := 0; i < 49; i++ {
			num := fmt.Sprintf("%02v", i+1)
			keys = append(keys, num)
			title = title + "," + num
		}
		return title + BreakLineTag
	}()

	// content
	for _, data := range op.Datas {
		csv = csv + data.Date.Format("2006-01-02")

		for _, key := range keys {
			hit := ""
			if key == data.Num_1 {
				hit = "1"
			}
			if key == data.Num_2 {
				hit = "1"
			}
			if key == data.Num_3 {
				hit = "1"
			}
			if key == data.Num_4 {
				hit = "1"
			}
			if key == data.Num_5 {
				hit = "1"
			}
			if key == data.Num_6 {
				hit = "1"
			}
			if key == data.NumSpecial {
				hit = "1"
			}
			csv = csv + "," + hit
		}
		csv = csv + BreakLineTag
	}

	//
	filename = fmt.Sprintf("nums_trending_%v.csv", op.Datas[len(op.Datas)-1].Date.Format("2006-01-02"))
	return
}
