package module

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_PlanA_GetCombinations(t *testing.T) {
	model.ConnMySQL()

	var all_hits = []lotto649op.Lotto649OPData{
		0: {SerialID: "test001", Num_1: "1", Num_2: "2", Num_3: "3", Num_4: "4", Num_5: "5", Num_6: "6", NumSpecial: "7"},
	}

	results := NewPlanA().GetCombinations(all_hits)
	for _, result := range results {
		fmt.Println(result)
	}
	/*
		output=>
		[1 2 3 4 5]
		[1 2 3 4 6]
		[1 2 3 4 7]
		[1 2 3 5 6]
		[1 2 3 5 7]
		[1 2 3 6 7]
		[1 2 4 5 6]
		[1 2 4 5 7]
		[1 2 4 6 7]
		[1 2 5 6 7]
		[1 3 4 5 6]
		[1 3 4 5 7]
		[1 3 4 6 7]
		[1 3 5 6 7]
		[1 4 5 6 7]
		[2 3 4 5 6]
		[2 3 4 5 7]
		[2 3 4 6 7]
		[2 3 5 6 7]
		[2 4 5 6 7]
		[3 4 5 6 7]
	*/
}

func Test_PlanA_1(t *testing.T) {
	combinations := [][]int{
		0: {1, 2, 3, 4, 5},
	}
	results, _ := NewPlanA().FillTo6(combinations)
	for _, result := range results {
		b, _ := json.MarshalIndent(result, "", "	")
		fmt.Println(string(b))
	}
}

func Test_PlanA_Run(t *testing.T) {
	model.ConnMySQL()
	raw_results, _ := model.NewLottery().FindAll()
	op := lotto649op.NewLotto649OP(raw_results)
	//
	plan_a := NewPlanA()
	filter_combinations := plan_a.GetCombinations(op.GetLotto649OPDatas())
	filter_combinations, _ = plan_a.FillTo6(filter_combinations)
	fmt.Println("總共=>", len(filter_combinations))

	//
	all_sets := map[string]struct{}{
		"4,11,20,25,32,39": {}, // 中獎號碼
		"4,11,20,30,32,39": {}, // 中獎號碼
		"1,2,4,5,11,13":    {},
	}
	all_sets = plan_a.RunFilter(all_sets, filter_combinations)
	fmt.Println(all_sets)
}

func Test_PlanA_real(t *testing.T) {
	model.ConnMySQL()
	raw_results, _ := model.NewLottery().FindAll()
	op := lotto649op.NewLotto649OP(raw_results)
	ori_hits := op.GetLotto649OPDatas()
	fmt.Println("hits =>", len(ori_hits))

	//
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, _ := time.ParseInLocation("2006-01-02", "1973-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-04", loc)
	hits := []lotto649op.Lotto649OPData{}
	for _, hit := range ori_hits {
		if hit.Date.Unix() < start.Unix() || hit.Date.Unix() > end.Unix() {
			continue
		}
		hits = append(hits, hit)
	}

	//

	plan_a := NewPlanA()
	filter_combinations := plan_a.GetCombinations(hits)
	filter_combinations, _ = plan_a.FillTo6(filter_combinations)
	fmt.Println("總共=>", len(filter_combinations))

	for _, filter_combination := range filter_combinations {
		sort.Ints(filter_combination)

		fmt.Println(filter_combination)
	}
}
