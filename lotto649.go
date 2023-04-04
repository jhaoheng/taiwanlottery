package lottery

/*
- 大樂透
- 目的: 讀取、寫入資料
*/

type ILotto649 interface {
}

type Lotto649 struct {
}

func NewLotto649() ILotto649 {
	return Lotto649{}
}

// 讀取, 透過 期數
func (lo *Lotto649) ReadBySerialID() []Lotto649 {
	return []Lotto649{}
}

// 將一筆資料寫入資料庫
func (lo *Lotto649) WriteDataInto() error {
	return nil
}
