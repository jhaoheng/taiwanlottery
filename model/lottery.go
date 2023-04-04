package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ILottery interface {
	SetID(id int64) *Lottery
	SetCategory(category LotteryCategory) *Lottery
	SetSerialID(SerialID string) *Lottery
	SetBallNumbers(ball_unmbers json.RawMessage) *Lottery
	/*
		- ex: 2006/01/02
		- 如果是民國(100/01/02), 則會轉換為西元
	*/
	SetDate(date string) *Lottery
	//
	Take() (Lottery, error)
	FindAll() ([]Lottery, error)
	Create() (Lottery, error)
	CreateInBatch(datas []Lottery) error
	Delete() error
}

type Lottery struct {
	db          *gorm.DB `gorm:"-"`
	ID          int64
	Category    LotteryCategory
	SerialID    string          // 期別, ex: 103000001
	BallNumbers json.RawMessage // ex: 1,2,3,4,5,6
	Date        time.Time       // 開獎日期, ex: 2006/01/01 15:04:05
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func NewLottery() ILottery {
	return &Lottery{
		db: db,
	}
}

func (Lottery) TableName() string {
	return "lottery"
}

func (model *Lottery) SetID(id int64) *Lottery {
	model.ID = id
	return model
}

func (model *Lottery) SetCategory(category LotteryCategory) *Lottery {
	model.Category = category
	return model
}

func (model *Lottery) SetSerialID(SerialID string) *Lottery {
	model.SerialID = SerialID
	return model
}

func (model *Lottery) SetBallNumbers(ball_unmbers json.RawMessage) *Lottery {
	model.BallNumbers = ball_unmbers
	return model
}

// date, ex: 2006/01/02, 如果是民國, 則會轉換為西元
func (model *Lottery) SetDate(date string) *Lottery {
	// 驗證日期
	date = func() string {
		tmp := strings.Split(date, "/")
		if len(tmp[0]) <= 3 {
			year := 1911
			// 判斷為民國年, ex: 100
			i, err := strconv.Atoi(tmp[0])
			if err != nil {
				panic(err)
			}
			year = i + year
			return fmt.Sprintf("%v/%v/%v", year, tmp[1], tmp[2])
		}
		return date
	}()
	loc, _ := time.LoadLocation("Asia/Taipei")
	tmp, err := time.ParseInLocation("2006/01/02", date, loc)
	if err != nil {
		panic(err)
	}
	model.Date = tmp
	return model
}

func (model *Lottery) Take() (Lottery, error) {
	output := Lottery{}
	tx := model.db.Where(model).Take(&output)
	return output, tx.Error
}

func (model *Lottery) FindAll() ([]Lottery, error) {
	output := []Lottery{}
	tx := model.db.Where(model).Find(&output)
	return output, tx.Error
}

func (model *Lottery) Create() (Lottery, error) {
	tx := model.db.Create(model)
	return *model, tx.Error
}

func (model *Lottery) CreateInBatch(datas []Lottery) error {
	tx := model.db.CreateInBatches(datas, 100)
	return tx.Error
}

func (model *Lottery) Update(vals Lottery) (Lottery, error) {
	tx := model.db.Table("lottery").Updates(vals)
	return *model, tx.Error
}

func (model *Lottery) Delete() error {
	tx := model.db.Delete(model)
	return tx.Error
}
