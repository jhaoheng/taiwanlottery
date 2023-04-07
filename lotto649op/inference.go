package lotto649op

/*
用來預測數字
*/

type InferenceResult struct {
	Num_1      string
	Num_2      string
	Num_3      string
	Num_4      string
	Num_5      string
	Num_6      string
	NumSpecial string
}

func (op *Lotto649OP) Inference() (result []InferenceResult) {

	return
}

/*
- 排除 歷史頭獎 取 5 個號碼, 再搭配兩個號碼
*/
func (r *InferenceResult) Excluded_1() {

}

/*
- 排除 上一次頭獎的 7 個號碼
*/
func (r *InferenceResult) Excluded_2() {

}
