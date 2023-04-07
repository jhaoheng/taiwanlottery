package lotto649op

import (
	"fmt"
	"testing"
	"time"
)

func Test_Excluded_1(t *testing.T) {

	// var sets PossibleSets = [][]int{
	// 	0: {11, 18, 20, 21, 35, 37},
	// 	1: {11, 18, 20, 21, 36, 38},
	// }

	op := NewLotto649OP(raw_results)
	sets := op.GetAllSets()
	fmt.Println("組數 =>", len(sets))
	new_set := sets[:len(sets)/100]
	fmt.Println("組數 =>", len(new_set))

	start := time.Now()

	//
	results := op.Excluded_1(new_set)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	fmt.Println("剩下組數 =>", len(results))

	fmt.Println(time.Until(start))
}
