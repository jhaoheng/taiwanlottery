package flowactions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/module"
	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/plan"
	"github.com/jhaoheng/taiwanlottery/lotto649op"
)

/*
	[目前策略] 指定時間區間
	1. 取得指定區間的中獎資料
	2. 取得推估的數據(消耗時間: 21.100764774s)
	3. 取得 plan_d 的組合 (消耗時間: 3.832917736s), 進行過濾
	4. 取得 plan_a 的數字（消耗時間: 1.160604563s），進行過濾
	5. 取得 plan_e 的數字（消耗時間: 53.774µs），進行過濾
	6. 進行 plan_f 的判斷


	=== 開始 GetAllSets ===
	消耗時間: 21.100764774s
	=== 開始 PlanD.GetConsecutiveSets() ===
	消耗時間: 22.598µs
	=== PlanD.FillTo6() ===
	消耗時間: 4.89455652s
	=== PlanA.GetCombinations() ===
	總共有: 21735, 執行時間: 7.085448ms
	=== PlanA.FillTo6() ===
	消耗時間: 976.342363ms
	=== PlanE.GetSpecificNums() ===
	消耗時間: 44.477µs
	最後剩下 => 2899068
*/

type IFlowA interface {
	GetNextHit() lotto649op.Lotto649OPData
	SetAllSetsMap(all_sets_map map[string]struct{}) *FlowA
	Run() *FlowA
	GetRankAndCSV() (csv string)
	GetRankAndExportOnlyHitIndexes() []plan.PlanFCountRankOnlyHitIndex
}

type FlowA struct {
	Loc            *time.Location
	Start          time.Time
	End            time.Time
	NextLastestHit lotto649op.Lotto649OPData   // end 時間的下一筆中獎資訊
	Hits           []lotto649op.Lotto649OPData // start - end 的中獎資訊
	//
	OriAllSetsMap map[string]struct{} // 原始組數
	AllSetsMap    map[string]struct{} // 計算結論
}

func NewFlowA(op lotto649op.ILotto649OP, start, end time.Time) IFlowA {
	// 1. 取得指定區間的中獎資料
	hits := []lotto649op.Lotto649OPData{}
	for _, hit := range op.GetLotto649OPDatas() {
		if hit.Date.Unix() < start.Unix() || hit.Date.Unix() > end.Unix() {
			continue
		}
		hits = append(hits, hit)
	}
	//
	next_lastest_hit := lotto649op.Lotto649OPData{
		SerialID: func() string {
			latest_sid := hits[len(hits)-1].SerialID
			i, _ := strconv.Atoi(latest_sid)
			return fmt.Sprintf("%d-預測", i+1)
		}(),
		Date:       time.Time{},
		Num_1:      "0",
		Num_2:      "0",
		Num_3:      "0",
		Num_4:      "0",
		Num_5:      "0",
		Num_6:      "0",
		NumSpecial: "0",
	}
	if len(op.GetNextDataByTime(end, 1)) != 0 {
		next_lastest_hit = op.GetNextDataByTime(end, 1)[0]
	}
	//
	fmt.Printf("全部 his 有: %v\n", len(hits))
	fmt.Printf("第一期數: %v, 最後一期期數: %v\n", hits[0].SerialID, hits[len(hits)-1].SerialID)
	return &FlowA{
		Start:          start,
		End:            end,
		NextLastestHit: next_lastest_hit,
		Hits:           hits,
	}
}

func (flow *FlowA) GetNextHit() lotto649op.Lotto649OPData {
	return flow.NextLastestHit
}

func (flow *FlowA) SetAllSetsMap(all_sets_map map[string]struct{}) *FlowA {
	fmt.Println("=== 設定 all_sets_map ===")
	flow.OriAllSetsMap = make(map[string]struct{})
	flow.OriAllSetsMap = all_sets_map
	return flow
}

/**/
func (flow *FlowA) Run() *FlowA {
	all_sets_map := make(map[string]struct{})
	if len(flow.OriAllSetsMap) == 0 {
		// 2. 取得推估的數據(消耗時間: 21.100764774s)
		_, all_sets_map = module.NewAllSets().Get()
	} else {
		for key := range flow.OriAllSetsMap {
			all_sets_map[key] = struct{}{}
		}
	}
	fmt.Println("原始組合 =>", len(all_sets_map))

	// 3. 取得 plan_d 的組合 (消耗時間: 3.832917736s), 排除連續數字, 進行過濾
	all_sets_map = func() map[string]struct{} {
		plan_d := plan.NewPlanD()
		filter_combinations := plan_d.GetConsecutiveSets(3)
		filter_combinations, _ = plan_d.FillTo6(filter_combinations)
		all_sets_map = plan_d.RunFilter(all_sets_map, filter_combinations)
		fmt.Println("剩下 =>", len(all_sets_map))
		return all_sets_map
	}()

	// 4. 取得 plan_a 的數字（消耗時間: 222.021544ms），進行過濾
	all_sets_map = func() map[string]struct{} {
		plan_a := plan.NewPlanA7Get5()
		filter_combinations := plan_a.GetCombinations(flow.Hits)
		filter_combinations, _ = plan_a.FillTo6(filter_combinations)
		all_sets_map = plan_a.RunFilter(all_sets_map, filter_combinations)
		fmt.Println("剩下 =>", len(all_sets_map))
		return all_sets_map
	}()

	flow.AllSetsMap = all_sets_map

	return flow
}

/**/
func (flow *FlowA) GetRankAndExportOnlyHitIndexes() []plan.PlanFCountRankOnlyHitIndex {
	plan_f := plan.NewPlanFCountRank(flow.AllSetsMap)
	ranks := plan_f.GetRank()
	return plan_f.ExportOnlyHitIndexes(ranks, flow.NextLastestHit)
}

/**/
func (flow *FlowA) GetRankAndCSV() (csv string) {
	plan_f := plan.NewPlanFCountRank(flow.AllSetsMap)
	ranks := plan_f.GetRank()
	return plan_f.ExportCSV(ranks, flow.NextLastestHit)
}
