package module

import (
	"fmt"
	"testing"
)

func Test_FillTo6(t *testing.T) {
	datas := [][]int{
		0: {1, 2, 3, 4},
	}
	results, _ := FillTo6(datas)
	for _, result := range results {
		fmt.Println(result)
	}
}
