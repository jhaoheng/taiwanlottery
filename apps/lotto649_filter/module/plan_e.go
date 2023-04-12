package module

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
)

/*
- 取得指定區間的中獎不重複數字
- 消耗時間: 53.774µs
*/
type IPlanE interface {
	GetSpecificNums(all_hits []lotto649op.Lotto649OPData, start, end time.Time) []int
	RunFilter(guess_sets map[string]struct{}, filter_nums []int) map[string]struct{}
}

type PlanE struct {
	Start time.Time
}

func NewPlanE() IPlanE {
	return &PlanE{
		Start: time.Now(),
	}
}

/*
- 取得指定區間的中獎不重複數字
*/
func (plan *PlanE) GetSpecificNums(all_hits []lotto649op.Lotto649OPData, start, end time.Time) []int {
	fmt.Println("=== PlanE.GetSpecificNums() ===")
	//
	map_num := map[string]struct{}{}
	for _, hit := range all_hits {
		if hit.Date.Unix() < start.Unix() || hit.Date.Unix() > end.Unix() {
			continue
		}
		if _, ok := map_num[hit.Num_1]; !ok {
			map_num[hit.Num_1] = struct{}{}
		}
		if _, ok := map_num[hit.Num_2]; !ok {
			map_num[hit.Num_2] = struct{}{}
		}
		if _, ok := map_num[hit.Num_3]; !ok {
			map_num[hit.Num_3] = struct{}{}
		}
		if _, ok := map_num[hit.Num_4]; !ok {
			map_num[hit.Num_4] = struct{}{}
		}
		if _, ok := map_num[hit.Num_5]; !ok {
			map_num[hit.Num_5] = struct{}{}
		}
		if _, ok := map_num[hit.Num_6]; !ok {
			map_num[hit.Num_6] = struct{}{}
		}
		if _, ok := map_num[hit.NumSpecial]; !ok {
			map_num[hit.NumSpecial] = struct{}{}
		}
	}

	results := []int{}
	for key := range map_num {
		i, _ := strconv.Atoi(key)
		results = append(results, i)
	}
	sort.Ints(results)
	//
	fmt.Printf("消耗時間: %v\n", -time.Until(plan.Start))
	return results
}

/*
 */
func (plan *PlanE) RunFilter(guess_sets map[string]struct{}, filter_nums []int) map[string]struct{} {

	keys := make([]string, 0, len(guess_sets))
	for k := range guess_sets {
		keys = append(keys, k)
	}

	//
	result_map := map[string]struct{}{}
	for _, key := range keys {
		is_found := false
		items := strings.Split(key, ",")
		for _, filter_num := range filter_nums {
			tmp := strconv.Itoa(filter_num)
			if items[0] == tmp {
				is_found = true
			}
			if items[1] == tmp {
				is_found = true
			}
			if items[2] == tmp {
				is_found = true
			}
			if items[3] == tmp {
				is_found = true
			}
			if items[4] == tmp {
				is_found = true
			}
			if items[5] == tmp {
				is_found = true
			}
			if is_found {
				break
			}
		}
		//
		if !is_found {
			if _, ok := result_map[key]; !ok {
				result_map[key] = struct{}{}
			}
		}
	}
	return result_map
}
