package main

import (
	"encoding/json"
	"fmt"
	"jhaoheng/taiwanlottery/fileop"
	"jhaoheng/taiwanlottery/model"
	"time"

	"gorm.io/gorm"
)

var filepaths = []string{
	"../../datas/2014/大樂透_2014.csv",
	"../../datas/2015/大樂透_2015.csv",
	"../../datas/2016/大樂透_2016.csv",
	"../../datas/2017/大樂透_2017.csv",
	"../../datas/2018/大樂透_201801_201812.csv",
	"../../datas/2019/大樂透_2019.csv",
	"../../datas/2020/大樂透_202001_202012.csv",
	"../../datas/2021/大樂透_2021.csv",
	"../../datas/2022/大樂透_2022.csv",
	// "../../datas/2014/大樂透加開獎項_2014.csv",
	// "../../datas/2015/大樂透加開獎項_2015.csv",
	// "../../datas/2016/大樂透加開獎項_2016.csv",
	// "../../datas/2017/大樂透加開獎項_2017.csv",
	// "../../datas/2018/大樂透加開獎項_201801_201812.csv",
	// "../../datas/2019/大樂透加開獎項_2019.csv",
	// "../../datas/2020/大樂透加開獎項_202001_202012.csv",
	// "../../datas/2021/大樂透加開獎項_2021.csv",
	// "../../datas/2022/大樂透加開獎項_2022.csv",
}

func main() {
	var csv_results []fileop.Lotto649CSV

	for _, filepath := range filepaths {
		file_op, err := fileop.NewFileOP().Read(filepath)
		if err != nil {
			panic(err)
		}
		tmps, err := file_op.ParsedLotto649CSV(",")
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
		lottery := model.Lottery{}
		if csv.GameName == "大樂透" {
			lottery = GetLotto649(csv, loc)
		} else if csv.GameName == "威力彩" {

		} else {
			panic("此 game 尚未設定")
		}
		lotterys = append(lotterys, lottery)
	}

	fmt.Printf("預計寫入資料庫有 %v 筆\n", len(lotterys))

	//
	model.ConnSQLite("../../sql.db", false)
	db_write_success := 0
	for _, lottery := range lotterys {
		obj, err := model.NewLottery().SetCategory(model.Lotto649).SetSerialID(lottery.SerialID).Take()
		if err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		if len(obj.SerialID) != 0 {
			continue
		}
		if _, err := lottery.Create(); err != nil {
			panic(err)
		}
		db_write_success++
	}

	fmt.Printf("寫入資料庫有 %v 筆\n", db_write_success)
}

/*
-
*/
func GetLotto649(csv fileop.Lotto649CSV, loc *time.Location) (lottery model.Lottery) {
	nums := model.Lotto649Nums{
		Num_1:       csv.Num_1,
		Num_2:       csv.Num_2,
		Num_3:       csv.Num_3,
		Num_4:       csv.Num_4,
		Num_5:       csv.Num_5,
		Num_6:       csv.Num_6,
		Num_special: csv.Num_special,
	}

	b, _ := json.Marshal(nums)

	return model.Lottery{
		Category:    model.Lotto649,
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
}

// /*
// -
// */
// func GetSuperlotto638(csv fileop, loc *time.Location) {

// }
