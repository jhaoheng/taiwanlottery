package plan

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test_RankCount_GetRank(t *testing.T) {
	hit_map := map[string]struct{}{
		"1,2,3,4,5,6": {},
		"2,3,4,5,6,7": {},
	}

	//
	ranks := NewPlanFCountRank(hit_map).GetRank()
	b, _ := json.MarshalIndent(ranks, "", "	")
	fmt.Println(string(b))
}

func Test_RankCount_ExportCSV(t *testing.T) {
	hit_map := map[string]struct{}{
		"1,2,3,4,5,6": {},
		"2,3,4,5,6,7": {},
	}

	//
	plan_f := NewPlanFCountRank(hit_map)
	ranks := plan_f.GetRank()
	csv := plan_f.ExportCSV(ranks, nil, nil)
	//
	filename := "./plan_f_export.csv"
	os.WriteFile(filename, []byte(csv), 0777)
}
