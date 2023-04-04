package model

type LotteryCategory string

const (
	Lotto649      LotteryCategory = "Lotto649"
	Superlotto638 LotteryCategory = "Superlotto638"
)

type Lotto649Nums struct {
	Num_1       string `json:"Num_1"`
	Num_2       string `json:"Num_2"`
	Num_3       string `json:"Num_3"`
	Num_4       string `json:"Num_4"`
	Num_5       string `json:"Num_5"`
	Num_6       string `json:"Num_6"`
	Num_special string `json:"Num_special"`
}

type Superlotto638Nums struct {
	Num_1              string `json:"Num_1"`
	Num_2              string `json:"Num_2"`
	Num_3              string `json:"Num_3"`
	Num_4              string `json:"Num_4"`
	Num_5              string `json:"Num_5"`
	Num_6              string `json:"Num_6"`
	Num_second_section string `json:"Num_second_section"`
}
