package module

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_PlanC_1(t *testing.T) {
	var all_hits = []lotto649op.Lotto649OPData{
		0: {SerialID: "test001", Num_1: "1", Num_2: "2", Num_3: "3", Num_4: "4", Num_5: "5", Num_6: "6", NumSpecial: "7"},
		1: {SerialID: "test002", Num_1: "1", Num_2: "3", Num_3: "5", Num_4: "7", Num_5: "9", Num_6: "11", NumSpecial: "13"},
	}

	results, exceptions := NewPlanC().GetconsecutiveHit(all_hits, 4)
	fmt.Println(results)
	fmt.Println(exceptions)
}

func Test_PlanC_2(t *testing.T) {

	model.ConnMySQL()

	raw_results, _ := model.NewLottery().FindAll()
	op := lotto649op.NewLotto649OP(raw_results)

	_, exceptions := NewPlanC().GetconsecutiveHit(op.GetLotto649OPDatas(), 4)
	for _, tmp := range exceptions {
		n1, _ := strconv.Atoi(tmp.Num_1)
		n2, _ := strconv.Atoi(tmp.Num_2)
		n3, _ := strconv.Atoi(tmp.Num_3)
		n4, _ := strconv.Atoi(tmp.Num_4)
		n5, _ := strconv.Atoi(tmp.Num_5)
		n6, _ := strconv.Atoi(tmp.Num_6)
		// n7, _ := strconv.Atoi(tmp.NumSpecial)
		numbers := []int{
			n1, n2, n3, n4, n5, n6,
		}
		sort.Ints(numbers)
		fmt.Printf("id: %v, date: %v, nums: %v, special: %v\n", tmp.SerialID, tmp.Date.Format("2006-01-02"), numbers, tmp.NumSpecial)
	}
	fmt.Println("總共 =>", len(exceptions))
}
