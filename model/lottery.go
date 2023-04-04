package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type ILottery interface {
	SetCategory(category string) *Lottery
	SetSerialID(SerialID string) *Lottery
	SetBallNumbers(ball_unmbers json.RawMessage) *Lottery
	SetDate(date string) *Lottery
	//
	Take() (Lottery, error)
	FindAll() ([]Lottery, error)
	Create() (Lottery, error)
}

type Lottery struct {
	db          *gorm.DB `gorm:"-"`
	ID          int64
	Category    string
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
	return "Lottery"
}

func (model *Lottery) SetCategory(category string) *Lottery {
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

// date, 2006/01/02
func (model *Lottery) SetDate(date string) *Lottery {
	tmp, err := time.Parse("2006/01/02", date)
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
