package lotto649op

import (
	"sort"
	"strconv"
	"time"
)

/*
- 取得區間內所有號碼出現的次數
*/
type NumCounts []NumCount

type NumCount struct {
	Num   string
	Count int
}

func (op *Lotto649OP) GetNumCount(start, end time.Time) (result NumCounts, start_data, end_data Lotto649OPData) {
	map_num := map[string]int{
		"01": 0, "02": 0, "03": 0, "04": 0, "05": 0, "06": 0, "07": 0, "08": 0, "09": 0, "10": 0,
		"11": 0, "12": 0, "13": 0, "14": 0, "15": 0, "16": 0, "17": 0, "18": 0, "19": 0, "20": 0,
		"21": 0, "22": 0, "23": 0, "24": 0, "25": 0, "26": 0, "27": 0, "28": 0, "29": 0, "30": 0,
		"31": 0, "32": 0, "33": 0, "34": 0, "35": 0, "36": 0, "37": 0, "38": 0, "39": 0, "40": 0,
		"41": 0, "42": 0, "43": 0, "44": 0, "45": 0, "46": 0, "47": 0, "48": 0, "49": 0,
	}

	start_data = Lotto649OPData{
		Date: time.Now().AddDate(100, 0, 0),
	}
	end_data = Lotto649OPData{
		Date: time.Now().AddDate(-100, 0, 0),
	}

	for _, data := range op.Datas {
		if data.Date.Unix() > start.Unix() && data.Date.Unix() < end.Unix() {
			map_num[data.Num_1]++
			map_num[data.Num_2]++
			map_num[data.Num_3]++
			map_num[data.Num_4]++
			map_num[data.Num_5]++
			map_num[data.Num_6]++
			map_num[data.NumSpecial]++

			//
			if start_data.Date.Unix() > data.Date.Unix() {
				start_data = data
			}
			if end_data.Date.Unix() < data.Date.Unix() {
				end_data = data
			}
		}
	}

	keys := make([]string, 0, len(map_num))
	for k := range map_num {
		keys = append(keys, k)
	}

	result = []NumCount{}
	for _, key := range keys {
		result = append(result, NumCount{
			Num:   key,
			Count: map_num[key],
		})
	}

	return result, start_data, end_data
}

func (data NumCounts) OrderNumCount() {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Count > data[j].Count
	})
}

func (data NumCounts) OrderByNum() {
	sort.Slice(data, func(i, j int) bool {
		a, _ := strconv.Atoi(data[i].Num)
		b, _ := strconv.Atoi(data[j].Num)
		return a < b
	})
}
