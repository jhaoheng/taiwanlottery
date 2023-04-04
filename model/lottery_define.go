package model

type LotteryCategory string

const (
	Lotto649      LotteryCategory = "大樂透"
	Superlotto638 LotteryCategory = "威力彩"
)

type Lotto649Nums struct {
	Num_1      string `json:"num_1"`
	Num_2      string `json:"num_2"`
	Num_3      string `json:"num_3"`
	Num_4      string `json:"num_4"`
	Num_5      string `json:"num_5"`
	Num_6      string `json:"num_6"`
	NumSpecial string `json:"num_special"`
}

type Superlotto638Nums struct {
	Num_1            string `json:"num_1"`
	Num_2            string `json:"num_2"`
	Num_3            string `json:"num_3"`
	Num_4            string `json:"num_4"`
	Num_5            string `json:"num_5"`
	Num_6            string `json:"num_6"`
	NumSecondSection string `json:"num_second_section"`
}
