package model

/*
- 過濾後的可能數組
*/

type ILotto649Filtered interface{}

type Lotto649Filtered struct {
}

func (Lotto649Filtered) TableName() string {
	return "lotto649_filtered"
}
