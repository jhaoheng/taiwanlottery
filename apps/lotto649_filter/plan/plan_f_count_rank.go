package plan

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
)

/*
[失敗]
- 似乎沒有一個穩定的趨勢
- 在排名中，儘管落於劣勢的號碼 or index，在下一次的中獎號碼中，依然有出現的可能性
- 所以透過每一次的數據變動，變更排行，查看趨勢，是不可能
*/

/*
[成功]
- 似乎可以看出某些趨勢
- 可排除的 index
	- min_index = 31,46,17
	- max_index = 1,22,30

每次數字出現在 index, 都會讓排名變動

# 排序
> 這可以看出來，雖然 index=30 總共出現 52 次中獎，但數字，集中在某些 index 上
> 平均出現次數數字是 38.285

index的中獎次數加總（index）
20(31)
27(46)
29 (17)
30(8), 30(25)
31(11), 31(19), 31(48)
32(9), 32(28)
33(2)
34(36)
34(3), 34(6), 34(24), 34(27)
35(44)
36(49)
37(21), 37(32)
38(4), 38(20), 38(42)
39(7), 39(15), 39(16), 39(26), 39(35), 39(45)
40(5), 40(10), 40(23), 40(41)
41(29), 41(43)
43(47)
44(12), 44(13), 44(37), 44(39)
45(34)
46(18), 46(40)
47(14), 47(33), 47(38)
45(1)
51(22)
52(30)
*/

/*
目的:
- 使用前，先排除掉，不可能中獎的排列
- 剩下的排列，個別計算數字，進行排名，查看趨勢
*/

type IPlanFCountRank interface {
	GetRank() []PlanFCountRankItem
	ExportCSV(datas []PlanFCountRankItem, compare_hit lotto649op.Lotto649OPData) (csv string)
	ExportOnlyHitIndexes(datas []PlanFCountRankItem, compare_hit lotto649op.Lotto649OPData) (results []PlanFCountRankOnlyHitIndex)
}

type PlanFCountRank struct {
	Start   time.Time
	DataMap map[string]struct{}
}

func NewPlanFCountRank(data_map map[string]struct{}) IPlanFCountRank {
	return &PlanFCountRank{
		Start:   time.Now(),
		DataMap: data_map,
	}
}

// 取得當前數字的出現排名
type PlanFCountRankItem struct {
	Num   int `json:"num"`
	Count int `json:"count"`
}

func (plan *PlanFCountRank) GetRank() []PlanFCountRankItem {
	fmt.Println("=== PlanE.GetRank ===")
	// init
	ranks := []PlanFCountRankItem{}
	rank_map := map[int]int{}
	for i := 1; i <= 49; i++ {
		rank_map[i] = 0
		ranks = append(ranks, PlanFCountRankItem{
			Num:   i,
			Count: 0,
		})
	}

	//
	for key := range plan.DataMap {
		nums := strings.Split(key, ",")
		if len(nums) != 6 {
			err := fmt.Errorf("數字數量錯誤, 必須是六, 但卻是: %v", key)
			panic(err)
		}
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		n3, _ := strconv.Atoi(nums[2])
		n4, _ := strconv.Atoi(nums[3])
		n5, _ := strconv.Atoi(nums[4])
		n6, _ := strconv.Atoi(nums[5])

		if _, ok := rank_map[n1]; ok {
			rank_map[n1]++
		}
		if _, ok := rank_map[n2]; ok {
			rank_map[n2]++
		}
		if _, ok := rank_map[n3]; ok {
			rank_map[n3]++
		}
		if _, ok := rank_map[n4]; ok {
			rank_map[n4]++
		}
		if _, ok := rank_map[n5]; ok {
			rank_map[n5]++
		}
		if _, ok := rank_map[n6]; ok {
			rank_map[n6]++
		}
	}

	//
	for key, value := range ranks {
		rank := PlanFCountRankItem{
			Num:   value.Num,
			Count: rank_map[value.Num],
		}
		ranks[key] = rank
	}

	// sort.Strings(rank_map)
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Count < ranks[j].Count
	})
	//
	fmt.Printf("消耗時間: %v\n", -time.Until(plan.Start))
	return ranks
}

func (plan *PlanFCountRank) ExportCSV(datas []PlanFCountRankItem, compare_hit lotto649op.Lotto649OPData) (csv string) {
	fmt.Printf("=== PlanE.ExportCSV ===\n")
	fmt.Printf("下一期參數: %+v\n", compare_hit)

	BreakLineTag := "\r\n"
	csv = ",num,count,hit" + BreakLineTag

	hit_map := func() map[int]struct{} {
		if compare_hit == (lotto649op.Lotto649OPData{}) {
			return map[int]struct{}{0: {}}
		}

		n1, _ := strconv.Atoi(compare_hit.Num_1)
		n2, _ := strconv.Atoi(compare_hit.Num_2)
		n3, _ := strconv.Atoi(compare_hit.Num_3)
		n4, _ := strconv.Atoi(compare_hit.Num_4)
		n5, _ := strconv.Atoi(compare_hit.Num_5)
		n6, _ := strconv.Atoi(compare_hit.Num_6)
		n7, _ := strconv.Atoi(compare_hit.NumSpecial)
		return map[int]struct{}{
			n1: {},
			n2: {},
			n3: {},
			n4: {},
			n5: {},
			n6: {},
			n7: {},
		}
	}()

	for index, data := range datas {
		num := data.Num
		count := data.Count
		is_hit := func() int {
			if _, ok := hit_map[num]; ok {
				return 1000000
			}
			return 0
		}()

		csv = csv + fmt.Sprintf("%d,%v,%v,%v", index+1, num, count, is_hit) + BreakLineTag
	}

	return
}

/*
y: index(order by count, ascending)
x: sid
z: hit

ex:
index--112000035
---33--------hit
*/
type PlanFCountRankOnlyHitIndex struct {
	Num         int
	Index       int
	Hit         bool
	HitSerialID string
}

func (plan *PlanFCountRank) ExportOnlyHitIndexes(datas []PlanFCountRankItem, compare_hit lotto649op.Lotto649OPData) (results []PlanFCountRankOnlyHitIndex) {
	fmt.Printf("=== PlanE.ExportOnlyHitIndexes ===\n")
	fmt.Printf("下一期參數: %+v\n", compare_hit)

	results = []PlanFCountRankOnlyHitIndex{}
	//
	hit_map := func() map[int]struct{} {
		if compare_hit == (lotto649op.Lotto649OPData{}) {
			return map[int]struct{}{0: {}}
		}

		n1, _ := strconv.Atoi(compare_hit.Num_1)
		n2, _ := strconv.Atoi(compare_hit.Num_2)
		n3, _ := strconv.Atoi(compare_hit.Num_3)
		n4, _ := strconv.Atoi(compare_hit.Num_4)
		n5, _ := strconv.Atoi(compare_hit.Num_5)
		n6, _ := strconv.Atoi(compare_hit.Num_6)
		n7, _ := strconv.Atoi(compare_hit.NumSpecial)
		return map[int]struct{}{
			n1: {},
			n2: {},
			n3: {},
			n4: {},
			n5: {},
			n6: {},
			n7: {},
		}
	}()
	for index, data := range datas {
		results = append(results, PlanFCountRankOnlyHitIndex{
			Num:         data.Num,
			Index:       index + 1,
			HitSerialID: compare_hit.SerialID,
			Hit: func() bool {
				if _, ok := hit_map[data.Num]; ok {
					return true
				}
				return false
			}(),
		})
	}

	return
}
