package plan

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jhaoheng/taiwanlottery/model"
	"gorm.io/gorm"
)

/*
- 計算 table num_index_hit 的報表
*/

type IPlanG interface {
	Get(table_name, sid string) (sums []model.NumIndexHitSum)
	ExportCSV(sums []model.NumIndexHitSum) (csv string)
}

type PlanG struct {
	SID               int
	SIDNumIndexHitMap map[int]int // 在 num_index_hit 中，hit 的 index
}

func NewPlanG() IPlanG {
	return &PlanG{}
}

// 取得指定 sid 的 hum_index_sum 資料
func (plan *PlanG) Get(table_name, sid string) (sums []model.NumIndexHitSum) {
	plan.SID, _ = strconv.Atoi(sid)
	if datas, err := model.NewNumIndexHit(table_name).SetSID(plan.SID).Take(); err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	} else {
		plan.SIDNumIndexHitMap = datas.ExportNumsToMap()
	}
	sums = []model.NumIndexHitSum{}
	for i := 1; i <= 49; i++ {
		sum, _ := model.NewNumIndexHit(table_name).Sum(plan.SID, i)
		sums = append(sums, sum)
	}

	/*
		- 進行分數的排列
		- 由小到大
	*/
	sort.Slice(sums, func(i, j int) bool {
		return sums[i].Total < sums[j].Total
	})
	return
}

func (plan *PlanG) ExportCSV(sums []model.NumIndexHitSum) (csv string) {
	type Result struct {
		Indexes []int
		Total   int
		Count   int
	}

	count_map := map[int]Result{}
	for _, sum := range sums {

		if _, ok := count_map[sum.Total]; !ok {
			count_map[sum.Total] = Result{
				Indexes: []int{sum.Index},
				Total:   sum.Total,
				Count:   1,
			}
		} else {
			indexes := count_map[sum.Total].Indexes
			indexes = append(indexes, sum.Index)
			count := count_map[sum.Total].Count + 1
			count_map[sum.Total] = Result{
				Indexes: indexes,
				Total:   sum.Total,
				Count:   count,
			}
		}
	}

	results := []Result{}
	for _, v := range count_map {
		results = append(results, v)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Total < results[j].Total
	})

	//
	BreakLineTag := "\r\n"
	csv = "indexes,sum,count,hit" + BreakLineTag
	for _, result := range results {
		// fmt.Printf("%+v\n", result)
		csv = csv + fmt.Sprintf("%v", result.Indexes) + ","
		csv = csv + fmt.Sprintf("%d", result.Total) + ","
		csv = csv + fmt.Sprintf("%d", result.Count) + ","

		text := ""
		for _, numindex := range result.Indexes {
			if plan.SIDNumIndexHitMap[numindex] == 1 {
				if len(text) != 0 {
					text = text + " "
				}
				text = text + fmt.Sprintf("%d", numindex)
			}
		}
		csv = csv + fmt.Sprintf("%v", text)

		csv = csv + BreakLineTag
	}

	return
}
