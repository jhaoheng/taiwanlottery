package module

import (
	"fmt"
	"testing"
)

func Test_PlanB(t *testing.T) {
	results := NewPlanB().GetCombinations()
	for _, result := range results {
		fmt.Println(result)
	}
}
