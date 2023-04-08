package module

import (
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

func Test_Filter_SearchLikeAndDel(t *testing.T) {
	model.IsDebug = true
	model.ConnMySQL()
	datas := [][]int{
		0: {1, 2, 3, 4, 5},
	}
	NewFilterTool().SearchLikeAndDel(datas)
}
