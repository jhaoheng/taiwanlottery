package module

import (
	"fmt"
	"testing"
)

func Test_PlanB(t *testing.T) {
	results := NewPlanB().GetCombinations(1, 49, 2, 6)
	for _, result := range results {
		fmt.Println(result)
	}

	/*
		output=>
		[1 2 3 4 5 6]
		[1 2 3 4 5 7]
		[1 3 4 5 6 7]
		[2 3 4 5 6 7]
	*/
}
