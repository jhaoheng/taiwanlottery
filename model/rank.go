package model

import (
	"encoding/json"
	"sort"

	"gorm.io/gorm"
)

type IRank interface {
	SetSID(sid int) IRank
	SetData(data []byte) IRank
	Take() (Rank, error)
	FindAll() ([]Rank, error)
	Create() (Rank, error)
}

type Rank struct {
	db   *gorm.DB `gorm:"-"`
	ID   int64
	SID  int `gorm:"column:sid"`
	Data string
}

type RankData struct {
	Num   int `json:"num"`
	Count int `json:"count"`
}

func (Rank) TableName() string {
	return "rank"
}

func NewRank() IRank {
	return &Rank{
		db: db,
	}
}

func (model *Rank) SetSID(sid int) IRank {
	model.SID = sid
	return model
}

func (model *Rank) SetData(data []byte) IRank {
	rans_datas := []RankData{}
	err := json.Unmarshal(data, &rans_datas)
	if err != nil {
		panic(err)
	}

	sort.Slice(rans_datas, func(i, j int) bool {
		return rans_datas[i].Count < rans_datas[j].Count
	})
	b, _ := json.Marshal(rans_datas)
	model.Data = string(b)
	return model
}

func (model *Rank) Take() (Rank, error) {
	output := Rank{}
	tx := model.db.Where(model).Take(&output)
	return output, tx.Error
}

func (model *Rank) FindAll() ([]Rank, error) {
	output := []Rank{}
	tx := model.db.Where(model).Find(&output)
	return output, tx.Error
}

func (model *Rank) Create() (Rank, error) {
	tx := model.db.Create(model)
	return *model, tx.Error
}
