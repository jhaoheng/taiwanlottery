package module

import (
	"fmt"
	"testing"
)

func Test_GetAllSets(t *testing.T) {

	/*
		[公式]
		- C(M) 取 N
		- ex: c7 取 6 => 7!/(6!*(7-6)!) = 7
		- ex: c8 取 6 => 8!/(6!*(8-6)!) = 28
	*/
	results, _ := NewAllSets().Get()
	fmt.Println(len(results))
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
}
