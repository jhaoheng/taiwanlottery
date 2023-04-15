package flow

import (
	"fmt"
	"time"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/module"
	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/plan"
	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
	"gorm.io/gorm"
)

/*
- flow_a 的延伸
- 從 flow a 取得的資料, 反覆進行計算
- 最後運算，產生 csv
*/

/*
[使用方法]
1.
*/

type IFlowB interface {
	Run(start_id int, op_times int, op lotto649op.ILotto649OP) (final_result [][]plan.PlanFCountRankOnlyHitIndex)
	CalIndexHitSum(multi_collections [][]plan.PlanFCountRankOnlyHitIndex) (results []FlowBCalIndexSum)
	ExportCSV(final_result [][]plan.PlanFCountRankOnlyHitIndex) (csv string)
}

type FlowB struct {
}

func NewFlowB() IFlowB {
	return &FlowB{}
}

/*
- start_id: 起始的 id，從此 id 取得樂透號碼後，開始從 flow_a 中，最後的可能數字中，進行運算
  - flow_a 會取得範圍數字，進行運算

- op_times: 預計執行的運算次數
*/
func (flow *FlowB) Run(start_id int, op_times int, op lotto649op.ILotto649OP) (final_result [][]plan.PlanFCountRankOnlyHitIndex) {

	final_result = [][]plan.PlanFCountRankOnlyHitIndex{}
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, err := time.ParseInLocation("2006-01-02", "1973-01-01", loc)
	if err != nil {
		panic(err)
	}

	_, all_sets_map := module.NewAllSets().Get()
	for i := start_id; i < start_id+op_times; i++ {
		data, err := model.NewLottery().SetID(int64(i)).Take()
		if err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		end := data.Date
		flow_a := NewFlowA(op, start, end).SetAllSetsMap(all_sets_map)
		results, _ := flow_a.Run().GetRankAndExportOnlyHitIndexes()
		final_result = append(final_result, results)
	}
	return
}

/*
- 利用 剩下的數量, 進行數字的排序 (index)
- 排序後，放入 hit 的資訊，若 hit 到，則 +1
- 取得每個 index 的總數
*/
type FlowBCalIndexSum struct {
	Sum   int
	Index int
}

func (flow *FlowB) CalIndexHitSum(multi_collections [][]plan.PlanFCountRankOnlyHitIndex) (results []FlowBCalIndexSum) {
	fmt.Println("=== flow_b.CalIndexHitSum() ===")
	nums := map[int]int{}
	for i := 1; i <= 49; i++ {
		nums[i] = 0
	}

	for _, collections := range multi_collections {
		for _, c := range collections {
			// fmt.Printf("%+v\n", c)
			if c.Hit {
				nums[c.Num] = nums[c.Num] + 1
			}
		}
	}

	results = []FlowBCalIndexSum{}
	collections := multi_collections[0]
	for _, c := range collections {
		results = append(results, FlowBCalIndexSum{
			Sum:   nums[c.Num],
			Index: c.Index,
		})
	}

	return
}

/**/
func (flow *FlowB) ExportCSV(final_result [][]plan.PlanFCountRankOnlyHitIndex) (csv string) {
	fmt.Println("=== 開始建立 csv ===")
	BreakLineTag := "\r\n"

	for i := 1; i <= 49; i++ {
		csv = csv + "," + fmt.Sprintf("%v", i)
	}
	csv = csv + BreakLineTag

	for i := 0; i < len(final_result); i++ {
		csv = csv + final_result[i][0].HitSerialID
		for j := 0; j < len(final_result[i]); j++ {
			value := ""
			if final_result[i][j].Hit {
				value = "1"
			}
			csv = csv + "," + fmt.Sprintf("%v", value)
		}
		csv = csv + BreakLineTag
	}
	return
}
