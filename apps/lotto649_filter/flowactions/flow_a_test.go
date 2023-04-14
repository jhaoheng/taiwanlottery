package flowactions

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/plan"
	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

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
	filename := plan_a.GetNextHit().SerialID + "-1.csv"
	os.WriteFile(filename, []byte(csv), 0777)

	fmt.Println("輸出完畢,", filename)
}

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
	result := plan_a.GetRankAndExportOnlyHitIndexes()
	//

	final_result := [][]plan.PlanFCountRankOnlyHitIndex{}
	final_result = append(final_result, result)

	/*
		建立 csv
	*/
	fmt.Println("=== 開始建立 csv ===")
	BreakLineTag := "\r\n"

	csv := ""

	for i := 1; i <= 49; i++ {
		csv = csv + "," + fmt.Sprintf("%v", i)
	}
	csv = csv + BreakLineTag

	for i := 0; i < len(final_result); i++ {
		csv = csv + final_result[i][0].HitSerialID
		for j := 0; j < len(final_result[i]); j++ {
			value := ""
			if final_result[i][j].Hit {
				value = "1"
			}
			csv = csv + "," + fmt.Sprintf("%v", value)
		}
		csv = csv + BreakLineTag
	}

	// fmt.Println(csv)
	filename := plan_a.GetNextHit().SerialID + "-2.csv"
	os.WriteFile(filename, []byte(csv), 0777)
}
