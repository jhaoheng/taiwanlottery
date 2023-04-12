package module

import (
	"fmt"
	"testing"
)

func Test_PlanD(t *testing.T) {
	N := 3
	plan_d := NewPlanD()
	filter_combinations := plan_d.GetConsecutiveSets(N)
	// for _, result := range filter_combinations {
	// 	fmt.Println(result)
	// }

	filter_combinations, _ = plan_d.FillTo6(filter_combinations)
	fmt.Println("總數 =>", len(filter_combinations))
	// for _, result := range filter_combinations {
	// 	fmt.Println(result)
	// }

	NewValidator().CheckRepeat(filter_combinations)

	//
	all_sets := map[string]struct{}{
		"1,2,3,4,5,6":   {},
		"1,2,4,5,11,13": {},
	}
	all_sets = plan_d.RunFilter(all_sets, filter_combinations)
	fmt.Println(all_sets)
}
