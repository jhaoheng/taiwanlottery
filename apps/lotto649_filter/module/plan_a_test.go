package module

import (
	"fmt"
	"testing"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_PlanA(t *testing.T) {
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
