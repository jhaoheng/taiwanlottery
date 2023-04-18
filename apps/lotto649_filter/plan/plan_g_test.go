package plan

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"testing"

	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_plang_run(t *testing.T) {
	model.ConnMySQL()
	sums := NewPlanG().Get("num_index_next_hit", "112000040", true)

	for _, sum := range sums {
		fmt.Printf("%+v\n", sum)
	}
}

func Test_plang_export(t *testing.T) {
	sid := "112000041"

	model.IsDebug = true
	model.ConnMySQL()
	sums := NewPlanG().Get("num_index_next_hit", "112000040", true)
	csv := NewPlanG().ExportCSV(sums)

	filename := fmt.Sprintf("plan_g_%v_tmp_2.csv", sid)
	os.WriteFile(filename, []byte(csv), 0777)
}

func Test_plang_GetWithCount(t *testing.T) {
	model.ConnMySQL()
	sid := 106000025
	sid_rows := 10   // 往後的數量
	data_count := -1 // 往前計算資料數量

	for i := sid; i < sid+sid_rows; i++ {
		sid_str := strconv.Itoa(i)
		plan_g := NewPlanG()
		sums := plan_g.GetWithCount("num_index_hit", sid_str, false, data_count)
		if len(sums) == 0 {
			return
		}
		csv := plan_g.ExportCSV(sums)
		filename := fmt.Sprintf("plan_g_test_%v_%d.csv", sid_str, data_count)
		os.WriteFile(filename, []byte(csv), 0777)
	}
}

func Test_plang_my(t *testing.T) {
	removes := []int{
		1, 2, 3, 9, 10, 11, 12, 13, 14, 16, 22, 23, 26, 27, 35, 42, 47, 49,
	}
	sort.Ints(removes)
	b, _ := json.Marshal(removes)
	fmt.Println(string(b))

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

	// // fmt.Println(datas)
	//
	//	for _, data := range datas {
	//		fmt.Printf("%v,", data)
	//	}
}
