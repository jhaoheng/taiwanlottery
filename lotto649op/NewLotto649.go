package lotto649op

import (
	"encoding/json"
	"time"

	"github.com/jhaoheng/taiwanlottery/model"
)

type ILotto649OP interface {
	// 取得區間內的數字次數
	GetNumCount(start, end time.Time) (total_count int, result NumCounts, start_data, end_data Lotto649OPData)
	// 取得最靠近此時間的下一次數據資料
	GetNextDataByTime(the_time time.Time) (the_data Lotto649OPData)
	// 取得累進範圍資料, 並取得最後的接下來 N 筆資料(用來比對)
	AccumulatedDatasByTime(start, end time.Time, future_count int) AccumulatedData
	// 取得時間範圍內, 以一期為單位, 重複的數字, 會有幾期
	GetRepeatNumRateInEachSID(repeat_num_count int, start, end time.Time)
	// 設定最少要相同的次數, 比對自己設定的號碼, 在歷史中, 正確率多高
	CheckCustomizedHits(hit_num_count int, nums ...string) (hit_lottery []Lotto649OPData)
	// 檢查歷史上, 中講過的號碼, 再次中獎機率
	CheckHitLotto(hit_num_count int) (results []CheckHitLottoResult)
	// 數字在時間上的趨勢
	ExportNumsTrending() (filename, csv string)
}

type Lotto649OP struct {
	Datas []Lotto649OPData
}

type Lotto649OPData struct {
	SerialID   string    // 期別, ex: 103000001
	Date       time.Time // 開獎日期, ex: 2006/01/01 15:04:05
	Num_1      string
	Num_2      string
	Num_3      string
	Num_4      string
	Num_5      string
	Num_6      string
	NumSpecial string
}

func NewLotto649OP(lotto649_raw_datas []model.Lottery) ILotto649OP {
	datas := []Lotto649OPData{}
	for _, raw_data := range lotto649_raw_datas {
		nums := model.Lotto649Nums{}
		json.Unmarshal(raw_data.BallNumbers, &nums)
		datas = append(datas, Lotto649OPData{
			SerialID:   raw_data.SerialID,
			Date:       raw_data.Date,
			Num_1:      nums.Num_1,
			Num_2:      nums.Num_2,
			Num_3:      nums.Num_3,
			Num_4:      nums.Num_4,
			Num_5:      nums.Num_5,
			Num_6:      nums.Num_6,
			NumSpecial: nums.NumSpecial,
		})
	}
	return &Lotto649OP{
		Datas: datas,
	}
}
