package module

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_Filter_DirectlyDel(t *testing.T) {
	model.IsDebug = true
	model.ConnMySQL()
	datas := [][]int{
		0: {01, 02, 03, 04, 05, 06},
		1: {01, 02, 03, 04, 05, 07},
	}
	NewFilterTool().DirectlyDel(datas)
}

func Test_Filter_SearchLikeAndDel_1(t *testing.T) {
	model.IsDebug = true
	model.ConnMySQL()
	datas := [][]int{
		0: {1, 2, 3, 4, 5},
		1: {2, 3, 4, 5, 6},
	}
	// NewFilterTool().SearchLikeAndDel(datas)
	NewFilterTool().SearchLikeAndDelWithGoroutine(datas)
}

func Test_Filter_SearchLikeAndDel_2(t *testing.T) {

	model.IsDebug = true
	model.ConnMySQL()
	N := 3
	datas := [][]int{}
	for i := 35; i <= 49-N; i++ {
		datas = append(datas, []int{i, i + 1, i + 2})
	}

	for _, data := range datas {
		fmt.Println(data)
	}

	NewFilterTool().SearchLikeAndDel(datas)
}

/*
-
*/
func Test_Filter_GetFilteredAndDoFilter(t *testing.T) {
	model.ConnMySQL()
	datas := [][]int{
		0: {1, 2, 4, 5, 7, 8},
	}
	results := NewFilterTool().GetFilteredAndDoFilter(datas)
	for _, result := range results {
		b, _ := json.MarshalIndent(result, "", "	")
		fmt.Println(string(b))
	}
}
