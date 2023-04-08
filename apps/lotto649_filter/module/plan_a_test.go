package module

import (
	"fmt"
	"testing"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_PlanA(t *testing.T) {
	model.ConnSQLite("../../../sql.db")

	var all_hits = []lotto649op.Lotto649OPData{
		0: {SerialID: "test001", Num_1: "1", Num_2: "2", Num_3: "3", Num_4: "4", Num_5: "5", Num_6: "6", NumSpecial: "7"},
	}

	results := NewPlanA().GetCombinations(all_hits)
	for _, result := range results {
		fmt.Println(result)
	}
}
