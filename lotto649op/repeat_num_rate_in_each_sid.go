package lotto649op

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
- 以每一期作為一個範本
- 重複出現相同範本的機率多高
*/

/*
- repeat_num_count: 每一期重複出現的號碼
*/
type RepeatNumRateInEachSID struct {
	CheckData Lotto649OPData
	Datas     []Lotto649OPData
}

func (op *Lotto649OP) GetRepeatNumRateInEachSID(repeat_num_count int, start, end time.Time) {

	// 過濾時間時間
	data_samples := []Lotto649OPData{}
	for _, data := range op.Datas {
		if data.Date.Unix() < start.Unix() || data.Date.Unix() > end.Unix() {
			continue
		}
		data_samples = append(data_samples, data)
	}

	//
	result := []RepeatNumRateInEachSID{}
	for _, xdata := range data_samples {
		//
		tmp := RepeatNumRateInEachSID{
			CheckData: xdata,
		}
		//
		for _, ydata := range data_samples {
			if xdata.SerialID == ydata.SerialID {
				continue
			}
			//
			repeats := []string{}
			map_x_data := map[string]int{
				xdata.Num_1:      0,
				xdata.Num_2:      0,
				xdata.Num_3:      0,
				xdata.Num_4:      0,
				xdata.Num_5:      0,
				xdata.Num_6:      0,
				xdata.NumSpecial: 0,
			}
			if _, ok := map_x_data[ydata.Num_1]; ok {
				repeats = append(repeats, ydata.Num_1)
			}
			if _, ok := map_x_data[ydata.Num_2]; ok {
				repeats = append(repeats, ydata.Num_2)
			}
			if _, ok := map_x_data[ydata.Num_3]; ok {
				repeats = append(repeats, ydata.Num_3)
			}
			if _, ok := map_x_data[ydata.Num_4]; ok {
				repeats = append(repeats, ydata.Num_4)
			}
			if _, ok := map_x_data[ydata.Num_5]; ok {
				repeats = append(repeats, ydata.Num_5)
			}
			if _, ok := map_x_data[ydata.Num_6]; ok {
				repeats = append(repeats, ydata.Num_6)
			}
			if _, ok := map_x_data[ydata.NumSpecial]; ok {
				repeats = append(repeats, ydata.NumSpecial)
			}

			//
			if len(repeats) >= repeat_num_count {
				tmp.Datas = append(tmp.Datas, ydata)
			}
		}
		//
		if len(tmp.Datas) != 0 {
			result = append(result, tmp)
		}
	}

	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))
	fmt.Printf("總共有 %v 組\n", len(result))
}
