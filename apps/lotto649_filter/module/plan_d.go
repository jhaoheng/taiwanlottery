package module

import (
	"fmt"
	"strconv"
	"time"
)

/*
- if N=3, 取得連續數字(N>=3)的組合
	- ex: 1,2,3,7,9,10
	- ex: 1,2,3,4,9,10

- N=3
消耗時間: 3.824118852s
總數 => 652827
*/

type IPlanD interface {
	GetConsecutiveSets(N int) [][]int
	FillTo6(combinations [][]int) (results [][]int, results_map map[string]struct{})
	RunFilter(guess_sets map[string]struct{}, filter_combinations [][]int) map[string]struct{}
}

type PlanD struct {
	Start time.Time
}

func NewPlanD() IPlanD {
	return &PlanD{
		Start: time.Now(),
	}
}

/*
 */
func (plan *PlanD) GetConsecutiveSets(N int) [][]int {
	fmt.Println("=== PlanD.GetConsecutiveSets() ===")
	datas := [][]int{}
	for i := 1; i <= 49-N; i++ {
		tmp := []int{}
		for j := 0; j < N; j++ {
			tmp = append(tmp, i+j)
		}
		datas = append(datas, tmp)
	}
	fmt.Printf("消耗時間: %v\n", -time.Until(plan.Start))
	return datas
}

/*
- 補齊六組號碼
*/
func (plan *PlanD) FillTo6(combinations [][]int) (results [][]int, results_map map[string]struct{}) {
	fmt.Println("=== PlanD.FillTo6() ===")
	//
	results, results_map = FillTo6(combinations)
	//
	fmt.Printf("消耗時間: %v\n", -time.Until(plan.Start))
	return results, results_map
}

/*
 */
func (plan *PlanD) RunFilter(guess_sets map[string]struct{}, filter_combinations [][]int) map[string]struct{} {
	for _, data := range filter_combinations {
		num_1 := strconv.Itoa(data[0])
		num_2 := strconv.Itoa(data[1])
		num_3 := strconv.Itoa(data[2])
		num_4 := strconv.Itoa(data[3])
		num_5 := strconv.Itoa(data[4])
		num_6 := strconv.Itoa(data[5])
		text := fmt.Sprintf("%v,%v,%v,%v,%v,%v", num_1, num_2, num_3, num_4, num_5, num_6)

		//
		delete(guess_sets, text)
	}
	return guess_sets
}
