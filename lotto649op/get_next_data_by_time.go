package lotto649op

import (
	"sort"
	"time"
)

func (op *Lotto649OP) GetNextDataByTime(the_time time.Time, next int) (results []Lotto649OPData) {
	results = []Lotto649OPData{}

	for _, data := range op.Datas {
		if data.Date.Unix()-the_time.Unix() > 0 {
			results = append(results, data)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Date.Unix() < results[j].Date.Unix()
	})
	return
}
