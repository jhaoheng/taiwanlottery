package module

import (
	"fmt"
	"sort"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

/*
- 會花比較久的時間, 第一次執行約一小時
- 只能排除掉約 一百萬筆
- 總共需要時間: 55 分鐘
*/

type IFilter_5_OfHitNums interface {
	GetNewSets(all_datas []lotto649op.Lotto649OPData, all_possible_sets lotto649op.PossibleSets) (new_sets lotto649op.PossibleSets)
	WriteToDB(new_sets lotto649op.PossibleSets)
}

type Filter_5_OfHitNums struct {
}

func New_Filter_5_OfHitNums() IFilter_5_OfHitNums {
	return &Filter_5_OfHitNums{}
}

/*
- 排除 => 歷史 頭獎 與 二獎 號碼 => 中五個號碼以上
- all_hits: 所有中獎資料
- all_possible_sets: 所有可能中獎的號碼
*/
func (filter *Filter_5_OfHitNums) GetNewSets(all_hits []lotto649op.Lotto649OPData, all_possible_sets lotto649op.PossibleSets) (new_sets lotto649op.PossibleSets) {
	fmt.Println("=== 開始過濾 ===")
	start := time.Now()
	type empty struct{}
	var member empty

	new_sets = [][]int{}

	// 創建一個 map，用來快速檢查 set 是否符合 op.Datas 的條件
	fmt.Println("準備資料中...")
	setMaps := make([]map[string]empty, len(all_possible_sets))
	for i, set := range all_possible_sets {
		setMaps[i] = make(map[string]empty)
		for _, num := range set {
			setMaps[i][fmt.Sprintf("%02d", num)] = member
		}
	}

	fmt.Println("開始...")
	for i, set := range all_possible_sets {
		if i%10000 == 0 {
			fmt.Printf("=== %v ===\n", i)
		}
		save_flag := true
		for _, data := range all_hits {
			hit := 0
			if _, ok := setMaps[i][data.Num_1]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_2]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_3]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_4]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_5]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.Num_6]; ok {
				hit++
			}
			if _, ok := setMaps[i][data.NumSpecial]; ok {
				hit++
			}
			if hit >= 5 {
				save_flag = false
				break
			}
		}
		if save_flag {
			new_sets = append(new_sets, set)
		}
	}

	fmt.Printf("所有組數: %v, 處理時間: %v\n", len(new_sets), -time.Until(start))
	return new_sets
}

/**/
func (filter *Filter_5_OfHitNums) WriteToDB(new_sets lotto649op.PossibleSets) {
	//
	// sets = sets[:len(sets)/10000]
	//
	start := time.Now()
	//
	fmt.Println("寫入資料格式化...")
	writes := []model.Lotto649AllSets{}
	for index, result := range new_sets {
		if index%1000 == 0 {
			fmt.Println("目前 =>", index)
		}
		writes = append(writes, model.Lotto649AllSets{
			Nums: func() string {
				if len(result) != 6 {
					panic("nums length not 6")
				}
				sort.Slice(result, func(i, j int) bool {
					return result[i] < result[j]
				})
				return fmt.Sprintf("%02d,%02d,%02d,%02d,%02d,%02d", result[0], result[1], result[2], result[3], result[4], result[5])
			}(),
		})
	}

	//
	fmt.Printf("格式化結束, 組數: %v, 處理時間: %v\n", len(writes), -time.Until(start))
	fmt.Printf("開始寫入資料庫\n")
	if err := model.NewLotto649AllSets().DeleteAll(); err != nil {
		panic(err)
	}
	if err := model.NewLotto649AllSets().CreateInBatch(writes, 1000); err != nil {
		panic(err)
	}
	fmt.Printf("done!!, 處理時間: %v\n\n", -time.Until(start))
}
