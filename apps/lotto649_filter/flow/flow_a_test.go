package flow

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

/*
- 欄位: index, num, count, next_sid_hit
- index: num 的 count 排序
- num: 數字 1-49
- count: 每個 num 可能出現的機率
- next_sid_hit: 下一期的中獎號碼
*/
func Test_flowa_GetRankAndCSV(t *testing.T) {
	model.ConnMySQL()
	raw_results, _ := model.NewLottery().FindAll()
	op := lotto649op.NewLotto649OP(raw_results)

	loc, _ := time.LoadLocation("Asia/Taipei")
	start, err := time.ParseInLocation("2006-01-02", "1973-01-01", loc)
	if err != nil {
		panic(err)
	}
	end, err := time.ParseInLocation("2006-01-02", "2023-04-07", loc)
	if err != nil {
		panic(err)
	}
	//
	plan_a := NewFlowA(op, start, end)
	csv := plan_a.Run().GetRankAndCSV()

	//
	filename := "flow_a_" + plan_a.GetNextHit().SerialID + "-1.csv"
	os.WriteFile(filename, []byte(csv), 0777)

	fmt.Println("輸出完畢,", filename)
}

/*
- 產生以指定 sid（sid）中獎號碼 在 index(累積數據後建立的 index) 的關係
- 設定 "2023-04-07" 會產生包含下一期的報表
*/
func Test_flowa_GetRankAndExportOnlyHitIndexes(t *testing.T) {
	model.ConnMySQL()
	raw_results, _ := model.NewLottery().FindAll()
	op := lotto649op.NewLotto649OP(raw_results)

	loc, _ := time.LoadLocation("Asia/Taipei")
	start, err := time.ParseInLocation("2006-01-02", "1973-01-01", loc)
	if err != nil {
		panic(err)
	}
	end, err := time.ParseInLocation("2006-01-02", "2023-04-07", loc)
	if err != nil {
		panic(err)
	}
	//
	plan_a := NewFlowA(op, start, end).Run()
	_, csv := plan_a.GetRankAndExportOnlyHitIndexes()

	// fmt.Println(csv)
	filename := "flow_a_2" + plan_a.GetNextHit().SerialID + "-2.csv"
	os.WriteFile(filename, []byte(csv), 0777)
}
