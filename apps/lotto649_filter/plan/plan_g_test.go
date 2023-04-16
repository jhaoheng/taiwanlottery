package plan

import (
	"fmt"
	"os"
	"testing"

	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_plang_run(t *testing.T) {
	model.ConnMySQL()
	sums := NewPlanG().Get("num_index_next_hit", "112000040")

	for _, sum := range sums {
		fmt.Printf("%+v\n", sum)
	}
}

func Test_plang_export(t *testing.T) {
	sid := "112000041"

	model.IsDebug = true
	model.ConnMySQL()
	sums := NewPlanG().Get("num_index_next_hit", sid)
	csv := NewPlanG().ExportCSV(sums)

	filename := fmt.Sprintf("plan_g_%v_tmp_2.csv", sid)
	os.WriteFile(filename, []byte(csv), 0777)
}

// func Test_plang_my(t *testing.T) {
// 	removes := []int{
// 		27, 47, 13, 9, 42, 10, 23, 11, 49, 26, 3, 14, 35, 1, 2, 16, 12, 22,
// 	}

// 	remove_map := map[int]struct{}{}
// 	for _, v := range removes {
// 		remove_map[v] = struct{}{}
// 	}

// 	//
// 	datas := []int{}
// 	for i := 1; i <= 49; i++ {
// 		if _, ok := remove_map[i]; !ok {
// 			datas = append(datas, i)
// 		}
// 	}

// 	// fmt.Println(datas)
// 	for _, data := range datas {
// 		fmt.Printf("%v,", data)
// 	}
// }
