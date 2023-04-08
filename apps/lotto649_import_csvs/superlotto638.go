package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jhaoheng/taiwanlottery/fileop"
	"github.com/jhaoheng/taiwanlottery/model"
)

func ImportSuperLotto638() {
	fmt.Println("=== 威力彩 ===")

	var filepaths = []string{
		"../../taiwan_lotto_csvs/2014/威力彩_2014.csv",
		"../../taiwan_lotto_csvs/2015/威力彩_2015.csv",
		"../../taiwan_lotto_csvs/2016/威力彩_2016.csv",
		"../../taiwan_lotto_csvs/2017/威力彩_2017.csv",
		"../../taiwan_lotto_csvs/2018/威力彩_201801_201812.csv",
		"../../taiwan_lotto_csvs/2019/威力彩_2019.csv",
		"../../taiwan_lotto_csvs/2020/威力彩_202001_202012.csv",
		"../../taiwan_lotto_csvs/2021/威力彩_2021.csv",
		"../../taiwan_lotto_csvs/2022/威力彩_2022.csv",
	}

	var csv_results []fileop.Superlotto638CSV

	for _, filepath := range filepaths {
		file_op, err := fileop.NewFileOP().Read(filepath)
		if err != nil {
			panic(err)
		}
		tmps, err := file_op.ParsedSuperlotto638(",")
		if err != nil {
			panic(err)
		}
		csv_results = append(csv_results, tmps...)
	}

	fmt.Printf("總共有 %v 筆\n", len(csv_results))

	//
	loc, _ := time.LoadLocation("Asia/Taipei")
	lotterys := []model.Lottery{}
	for _, csv := range csv_results {
		if csv.GameName != "威力彩" {
			panic("game name 錯誤")
		}
		lottery := func(csv fileop.Superlotto638CSV, loc *time.Location) (lottery model.Lottery) {
			nums := model.Superlotto638Nums{
				Num_1:            csv.Num_1,
				Num_2:            csv.Num_2,
				Num_3:            csv.Num_3,
				Num_4:            csv.Num_4,
				Num_5:            csv.Num_5,
				Num_6:            csv.Num_6,
				NumSecondSection: csv.Num_second_section,
			}

			b, _ := json.Marshal(nums)

			return model.Lottery{
				Category:    model.Superlotto638,
				SerialID:    csv.SerialID,
				BallNumbers: b,
				Date: func() time.Time {
					tmp, err := time.ParseInLocation("2006/01/02", csv.Date, loc)
					if err != nil {
						panic(err)
					}
					return tmp
				}(),
			}
		}(csv, loc)

		lotterys = append(lotterys, lottery)
	}

	fmt.Printf("預計寫入資料庫有 %v 筆\n", len(lotterys))

	//
	WriteToDB(model.Superlotto638, lotterys)
}
