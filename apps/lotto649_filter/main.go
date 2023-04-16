package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/flow"
	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/module"
	"github.com/jhaoheng/taiwanlottery/model"
)

/*
	raw_results, _ = model.NewLottery().FindAll()
	op = lotto649op.NewLotto649OP(raw_results)
	module.GetAllPossiblility(op) // 取得所有組數, 並寫入資料庫
*/

func main() {
	// model.IsDebug = true
	err := model.ConnMySQL()
	if err != nil {
		panic(err)
	}
	SpecificRun(112000042)
}

/*
- 執行指定 serial_id
- 指定最後一個 serial_id 則是預測下一輪
*/
func SpecificRun(inferential_sid int) {
	//
	_, all_sets_map := module.NewAllSets().Get()
	//
	start := time.Now()
	flow_c := flow.NewFlowC(inferential_sid, all_sets_map)
	// flow_c.RunPlansAndGetReports()
	result := flow_c.RunPlansAndGetReports()
	flow_c.SaveReports(result, "預測_"+strconv.Itoa(inferential_sid))

	fmt.Println(-time.Until(start))
	fmt.Printf("\n\n\n\n\n")
}

/*
- 執行多次
- 用於一次產生多個報表
- 若 db 沒資料，則新增資料到 num_index_hit 以及 num_index_next_hit
*/
func MultiRun() {
	_, all_sets_map := module.NewAllSets().Get()

	for id := 1018; id < 1039; id++ {
		start := time.Now()
		item, err := model.NewLottery().SetID(int64(id)).Take()
		if err != nil {
			panic(err)
		}
		inferential_sid, _ := strconv.Atoi(item.SerialID)
		flow_c := flow.NewFlowC(inferential_sid, all_sets_map)
		// flow_c.RunPlansAndGetReports()
		result := flow_c.RunPlansAndGetReports()
		flow_c.SaveReports(result, "預測_"+strconv.Itoa(inferential_sid))

		fmt.Println(-time.Until(start))
		fmt.Printf("\n\n\n\n\n")
	}
}
