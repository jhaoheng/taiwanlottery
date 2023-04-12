package module

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_PlanE_Run(t *testing.T) {

	model.ConnMySQL()
	raw_results, _ := model.NewLottery().FindAll()
	op := lotto649op.NewLotto649OP(raw_results)

	loc, _ := time.LoadLocation("Asia/Taipei")
	start, _ := time.ParseInLocation("2006-01-02", "2023-04-07", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-07", loc)

	//
	plan_e := NewPlanE()
	filter_combinations := plan_e.GetSpecificNums(op.GetLotto649OPDatas(), start, end)
	fmt.Println("取得該過濾的數字 =>", filter_combinations)

	//
	all_sets := map[string]struct{}{
		"1,2,3,5,6,7":       {},
		"1,2,4,5,11,13":     {},
		"11,13,20,33,42,45": {},
	}
	all_sets = plan_e.RunFilter(all_sets, filter_combinations)
	fmt.Println(func() string {
		b, _ := json.MarshalIndent(all_sets, "", "	")
		return string(b)
	}())

	/*
		=== PlanE.GetSpecificNums() ===
		消耗時間: 39.137µs
		取得該過濾的數字 => [4 11 20 25 30 32 39]
		{
			"1,2,3,5,6,7": {}
		}
	*/
}
