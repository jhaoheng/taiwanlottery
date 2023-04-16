package flow

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/module"
	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/plan"
	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
	"gorm.io/gorm"
)

/*
- 使用 flow_c 需先確認資料庫資料是否齊全，否則預測出來的報表會有失準機會
*/

/*
- 如果要預測 112000041
		- 取得基本的兩份報表後
		- 計算 num_index_sum, 並取得趨勢分析表
		- 手動判斷出 112000041 不可能出現的數字
*/

type IFlowC interface {
	RunPlansAndGetReports() []*CSVData
	// folderName, ex: 112000042
	SaveReports(csvs []*CSVData, folderName string)
}

type FlowC struct {
	op             lotto649op.ILotto649OP
	InferentialSID int                         // 要推論的 sid
	FirstSID       int                         // 第一筆資料的 sid
	LastSID        int                         // 最後一筆資料的 sid
	AllHits        []lotto649op.Lotto649OPData // 所有中獎號碼
	AllSetsMap     map[string]struct{}         // 所有排列組合, ex: ["1,2,3,4,5,6"]=struct{}{}
}

// inferential_sid: 要預測的 sid
func NewFlowC(inferential_sid int) IFlowC {
	first_sid := 103000001
	last_sid := inferential_sid - 1
	lotterys, err := model.NewLottery().FindBetweenSID(strconv.Itoa(first_sid), strconv.Itoa(last_sid))
	if err != nil {
		panic(err)
	}

	//
	op := lotto649op.NewLotto649OP(lotterys)
	op.GetLotto649OPDatas()
	all_hits := op.GetLotto649OPDatas()
	//
	fmt.Printf("全部 his 有: %v\n", len(all_hits))
	fmt.Printf("第一期期數: %v, 最後一期期數: %v\n", all_hits[0].SerialID, all_hits[len(all_hits)-1].SerialID)
	fmt.Printf("預測 sid: %d\n", inferential_sid)
	fmt.Println()

	/*
		修正正確的 last_sid, 因為有可能 last_sid = 104000001 - 1 = 104000000, 此為錯誤的 sid
	*/
	last_sid = func() int {
		i, _ := strconv.Atoi(all_hits[len(all_hits)-1].SerialID)
		return i
	}()

	//
	_, all_sets_map := module.NewAllSets().Get()
	fmt.Println("all_sets_map 一開始有:", len(all_sets_map))
	return &FlowC{
		op:             op,
		InferentialSID: inferential_sid,
		FirstSID:       first_sid,
		LastSID:        last_sid,
		AllHits:        all_hits,
		AllSetsMap:     all_sets_map,
	}
}

type CSVData struct {
	Filename string
	CSV      string
}

/*
[說明]
- 執行 plans, 過濾掉可能的組合(all_sets_map)
- 取得兩份報表
 1. 以當前可能的數字資料, 輸出 index,num,count 報表
 2. 以當前可能的數字資料, 輸出 最後一期中獎號碼的數字落點 報表
*/
func (flow *FlowC) RunPlansAndGetReports() []*CSVData {

	//
	all_sets_map := flow.AllSetsMap

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
		filter_combinations := plan_a.GetCombinations(flow.AllHits)
		filter_combinations, _ = plan_a.FillTo6(filter_combinations)
		all_sets_map = plan_a.RunFilter(all_sets_map, filter_combinations)
		fmt.Println("剩下 =>", len(all_sets_map))
		return all_sets_map
	}()

	//
	/*
		執行 plan_f 取得報表
	*/
	csv_data_1 := flow.GetReport_1(all_sets_map)
	csv_data_2 := flow.GetReport_2(all_sets_map)
	csv_data_3 := flow.GetReport_3(all_sets_map)

	/*
		- 取得 plan_g 報表
	*/
	// 從 table::num_index_hit 中取得
	plan_g_from_csv_data_2 := func() *CSVData {
		plan_g := plan.NewPlanG()
		sums := plan_g.Get("num_index_hit", strconv.Itoa(flow.LastSID))
		csv := plan_g.ExportCSV(sums)

		return &CSVData{
			Filename: fmt.Sprintf("%v_plan_g_趨勢圖.csv", flow.LastSID),
			CSV:      csv,
		}
	}()
	// 從 table::num_index_next_hit 中取得
	plan_g_from_csv_data_3 := func() *CSVData {
		plan_g := plan.NewPlanG()
		sums := plan_g.Get("num_index_next_hit", strconv.Itoa(flow.InferentialSID))
		csv := plan_g.ExportCSV(sums)

		return &CSVData{
			Filename: fmt.Sprintf("%v_plan_g_趨勢圖.csv", flow.InferentialSID),
			CSV:      csv,
		}
	}()
	//
	csv_datas := []*CSVData{}
	csv_datas = append(csv_datas, csv_data_1, csv_data_2, csv_data_3, plan_g_from_csv_data_2, plan_g_from_csv_data_3)
	return csv_datas
}

/*
- 取得報表 1
*/
func (flow *FlowC) GetReport_1(all_sets_map map[string]struct{}) *CSVData {
	//
	plan_f := plan.NewPlanFCountRank(all_sets_map)
	ranks := plan_f.GetRank()
	//
	csv_data_1 := CSVData{
		Filename: fmt.Sprintf("%v_plan_f_nums_count累積_排序.csv", flow.LastSID),
		CSV: plan_f.ExportCSV(
			ranks,
			&flow.AllHits[len(flow.AllHits)-1],
			func() *lotto649op.Lotto649OPData {
				// 檢查 inferential_sid 是否存在資料庫
				datas, err := model.NewLottery().SetSerialID(strconv.Itoa(flow.InferentialSID)).FindAll()
				if err != nil {
					panic(err)
				}
				if len(datas) == 0 {
					return nil
				}
				return &lotto649op.NewLotto649OP(datas).GetLotto649OPDatas()[0]
			}(),
		),
	}
	return &csv_data_1
}

/*
- 取得報表 2, 最後一期中獎號碼與累積數據的關係
- 寫入 num_index_hit
*/
func (flow *FlowC) GetReport_2(all_sets_map map[string]struct{}) *CSVData {
	//
	plan_f := plan.NewPlanFCountRank(all_sets_map)
	ranks := plan_f.GetRank()
	//
	last_hit := flow.AllHits[len(flow.AllHits)-1]
	csv_data_2 := CSVData{
		Filename: fmt.Sprintf("%v_plan_f_hit_落點.csv", last_hit.SerialID),
		CSV: func() string {
			results := plan_f.ExportOnlyHitIndexes(ranks, &last_hit)

			// 將 data 寫入 table::num_index_hit, 如果資料不存在
			sid_int, _ := strconv.Atoi(last_hit.SerialID)
			if _, err := model.NewNumIndexHit("num_index_hit").SetSID(sid_int).Take(); err != nil && err != gorm.ErrRecordNotFound {
				panic(err)
			} else if err == gorm.ErrRecordNotFound {
				num_indexes := []model.NumIndex{}
				for _, result := range results {
					num_indexes = append(num_indexes, model.NumIndex{
						Index: result.Index,
						Hit: func() int {
							if result.Hit {
								return 1
							}
							return 0
						}(),
					})
				}
				if _, err := model.NewNumIndexHit("num_index_hit").SetSID(sid_int).SetNumIndexes(num_indexes).Create(); err != nil {
					panic(err)
				}
			}

			//建立這一次檔案的 csv
			csv := ""
			BreakLineTag := "\r\n"

			for i := 1; i <= 49; i++ {
				csv = csv + "," + fmt.Sprintf("%v", i)
			}
			csv = csv + BreakLineTag

			// //
			// b, _ := json.MarshalIndent(results, "", "	")
			// fmt.Println(string(b))
			csv = csv + results[0].HitSerialID
			for _, result := range results {
				value := ""
				if result.Hit {
					value = "1"
				}
				csv = csv + "," + fmt.Sprintf("%v", value)
			}
			csv = csv + BreakLineTag
			return csv
		}(),
	}
	return &csv_data_2
}

/*
- 取得報表 3, 下一期中獎號碼與累積數據的關係
- 如果沒有下一期中獎號碼，則不運算
- 寫入 num_index_next_hit
*/
func (flow *FlowC) GetReport_3(all_sets_map map[string]struct{}) *CSVData {
	//
	plan_f := plan.NewPlanFCountRank(all_sets_map)
	ranks := plan_f.GetRank()
	// 取得下一期中獎號碼, 如果沒有則不運算
	next_hit := func() *lotto649op.Lotto649OPData {
		// 檢查 inferential_sid 是否存在資料庫
		datas, err := model.NewLottery().SetSerialID(strconv.Itoa(flow.InferentialSID)).FindAll()
		if err != nil {
			panic(err)
		}
		if len(datas) == 0 {
			return nil
		}
		return &lotto649op.NewLotto649OP(datas).GetLotto649OPDatas()[0]
	}()

	if next_hit == nil {
		return nil
	}

	csv_data_3 := CSVData{
		Filename: fmt.Sprintf("%v_plan_f_hit_落點.csv", next_hit.SerialID),
		CSV: func() string {
			results := plan_f.ExportOnlyHitIndexes(ranks, next_hit)

			// 將 data 寫入 table::num_index_next_hit, 如果資料不存在
			sid_int, _ := strconv.Atoi(next_hit.SerialID)
			if _, err := model.NewNumIndexHit("num_index_next_hit").SetSID(sid_int).Take(); err != nil && err != gorm.ErrRecordNotFound {
				panic(err)
			} else if err == gorm.ErrRecordNotFound {
				num_indexes := []model.NumIndex{}
				for _, result := range results {
					num_indexes = append(num_indexes, model.NumIndex{
						Index: result.Index,
						Hit: func() int {
							if result.Hit {
								return 1
							}
							return 0
						}(),
					})
				}
				if _, err := model.NewNumIndexHit("num_index_next_hit").SetSID(sid_int).SetNumIndexes(num_indexes).Create(); err != nil {
					panic(err)
				}
			}

			//建立這一次檔案的 csv
			csv := ""
			BreakLineTag := "\r\n"

			for i := 1; i <= 49; i++ {
				csv = csv + "," + fmt.Sprintf("%v", i)
			}
			csv = csv + BreakLineTag

			// //
			// b, _ := json.MarshalIndent(results, "", "	")
			// fmt.Println(string(b))
			csv = csv + results[0].HitSerialID
			for _, result := range results {
				value := ""
				if result.Hit {
					value = "1"
				}
				csv = csv + "," + fmt.Sprintf("%v", value)
			}
			csv = csv + BreakLineTag
			return csv
		}(),
	}
	return &csv_data_3
}

// -
func (flow *FlowC) SaveReports(csvs []*CSVData, folderName string) {
	//
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		// 資料夾不存在，創建資料夾
		if err := os.Mkdir(folderName, 0755); err != nil {
			fmt.Printf("Failed to create folder '%s': %s\n", folderName, err)
		} else {
			fmt.Printf("Folder '%s' created successfully.\n", folderName)
		}
	} else {
		fmt.Printf("Folder '%s' exists.\n", folderName)
	}
	//

	for _, csv_data := range csvs {
		if csv_data == nil {
			continue
		}
		filename := time.Now().Format("20060102150405") + "-" + csv_data.Filename
		os.WriteFile(folderName+"/"+filename, []byte(csv_data.CSV), 0777)
	}
}
