package lotto649op

import (
	"fmt"
	"time"
)

/*
[累進制-區間資料]
- 從第一年的資料開始計算次數
- 每次計算完畢, 取得最後一筆資料(sid)的後十筆資料
*/

type AccumulatedData struct {
	TotalCount  int
	NumCounts   NumCounts
	StartData   Lotto649OPData
	EndData     Lotto649OPData
	FutureDatas []Lotto649OPData
}

func (op *Lotto649OP) AccumulatedDatasByTime(start, end time.Time, future_count int) AccumulatedData {
	total_count, num_counts, start_data, end_data := op.GetNumCount(start, end)
	num_counts.OrderNumCount()

	future_datas := []Lotto649OPData{}
	end_date := end_data.Date
	for i := 0; i < future_count; i++ {
		datas := op.GetNextDataByTime(end_date, 1)
		if len(datas[0].NumSpecial) == 0 {
			break
		}
		future_datas = append(future_datas, datas[0])
		// b, _ := json.MarshalIndent(data, "", "	")
		// fmt.Println(string(b))
		end_date = datas[0].Date
	}

	return AccumulatedData{
		TotalCount:  total_count,
		NumCounts:   num_counts,
		StartData:   start_data,
		EndData:     end_data,
		FutureDatas: future_datas,
	}
}

/*
X: count, percentage, 指定期數的中獎號碼, 每一期的中獎數字
Y: 號碼
*/
func (accumulated_data *AccumulatedData) ExportCSV_1() (filename, data string) {

	// sid_hit_num: 期號中獎號碼
	type SidHitNum map[string]bool
	future_data_keys := []string{}
	future_data_map := map[string]SidHitNum{}
	for _, future_data := range accumulated_data.FutureDatas {
		future_data_keys = append(future_data_keys, future_data.SerialID)
		future_data_map[future_data.SerialID] = SidHitNum{
			"01": false, "02": false, "03": false, "04": false, "05": false, "06": false, "07": false, "08": false, "09": false, "10": false,
			"11": false, "12": false, "13": false, "14": false, "15": false, "16": false, "17": false, "18": false, "19": false, "20": false,
			"21": false, "22": false, "23": false, "24": false, "25": false, "26": false, "27": false, "28": false, "29": false, "30": false,
			"31": false, "32": false, "33": false, "34": false, "35": false, "36": false, "37": false, "38": false, "39": false, "40": false,
			"41": false, "42": false, "43": false, "44": false, "45": false, "46": false, "47": false, "48": false, "49": false,
		}
		//
		future_data_map[future_data.SerialID][future_data.Num_1] = true
		future_data_map[future_data.SerialID][future_data.Num_2] = true
		future_data_map[future_data.SerialID][future_data.Num_3] = true
		future_data_map[future_data.SerialID][future_data.Num_4] = true
		future_data_map[future_data.SerialID][future_data.Num_5] = true
		future_data_map[future_data.SerialID][future_data.Num_6] = true
		future_data_map[future_data.SerialID][future_data.NumSpecial] = true
	}

	//
	BreakLineTag := "\r\n"
	var csv string

	csv = func() string {
		// 建立 title
		title := "num, count, percentage"
		for _, key := range future_data_keys {
			title = title + " ," + key
		}
		return title + BreakLineTag
	}()

	// 建立內容
	accumulated_data.NumCounts.OrderByNum()
	for _, num_count := range accumulated_data.NumCounts {
		//
		csv = csv + num_count.Num + ","                      // 第一欄 = 數字
		csv = csv + fmt.Sprintf("%v", num_count.Count) + "," // 第二欄 = 數量
		csv = csv + fmt.Sprintf("%v", num_count.Percentage)  // 第三欄 = 百分比

		//
		for i := 0; i < len(future_data_keys); i++ {
			csv = csv + func() string {
				key := future_data_keys[i]
				content := ","
				if future_data_map[key][num_count.Num] {
					content = ",ok"
				}
				// fmt.Println(key, content)
				if i == len(future_data_keys)-1 {
					return content
				}
				return content
			}()
		}

		csv = csv + BreakLineTag
	}

	// 設定檔案名稱
	filename = fmt.Sprintf("accumulated,%v-%v,%v.csv", accumulated_data.StartData.SerialID, accumulated_data.EndData.SerialID, len(accumulated_data.FutureDatas))

	return filename, csv
}

// /*
// X: 日期, 2006-01-02
// Y: 數字
// */
// func (accumulated_data *AccumulatedData) ExportCSV_2() (filename, data string) {
// 	return
// }
