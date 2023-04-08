package model

import (
	"fmt"
	"sort"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
- 所有可能的數組
*/

type ILotto649AllSets interface {
	SetID(id int64) *Lotto649AllSets
	SetNums(nums []int) *Lotto649AllSets
	//
	Take() (Lotto649AllSets, error)
	FindAll() ([]Lotto649AllSets, error)
	// FindNumsLike(texts []string) ([]Lotto649AllSets, error)
	Create() (Lotto649AllSets, error)
	CreateInBatch(datas []Lotto649AllSets, batch_size int) error
	Update(vals Lotto649AllSets) (Lotto649AllSets, error)
	Delete() error
	DeleteAll() error
	OrderByDESC(item_name string) *Lotto649AllSets
}

type Lotto649AllSets struct {
	db        *gorm.DB  `gorm:"-"`
	ID        int64     `gorm:"primaryKey"` //
	Nums      string    // sort ascending and only 6 nums, ex: 1,2,3,4,5,6
	UpdatedAt time.Time //
	CreatedAt time.Time //
}

func (Lotto649AllSets) TableName() string {
	return "lotto649_all_sets"
}

func NewLotto649AllSets() ILotto649AllSets {
	return &Lotto649AllSets{
		db: db,
	}
}

func (model *Lotto649AllSets) SetID(id int64) *Lotto649AllSets {
	model.ID = id
	return model
}

func (model *Lotto649AllSets) SetNums(nums []int) *Lotto649AllSets {
	if len(nums) != 6 {
		panic("nums length not 6")
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	model.Nums = fmt.Sprintf("%02d,%02d,%02d,%02d,%02d,%02d", nums[0], nums[1], nums[2], nums[3], nums[4], nums[5])
	return model
}

func (model *Lotto649AllSets) Take() (Lotto649AllSets, error) {
	output := Lotto649AllSets{}
	tx := model.db.Where(model).Take(&output)
	return output, tx.Error
}

func (model *Lotto649AllSets) FindAll() ([]Lotto649AllSets, error) {
	output := []Lotto649AllSets{}
	tx := model.db.Where(model).Find(&output)
	return output, tx.Error
}

// // text, ex: %abc%
// func (model *Lotto649AllSets) FindNumsLike(texts []string) ([]Lotto649AllSets, error) {
// 	output := []Lotto649AllSets{}
// 	for _, text := range texts {

// 	}
// 	tx := model.db.Where(model).Where("nums LIKE ?", text).Find(&output)
// 	return output, tx.Error
// }

func (model *Lotto649AllSets) Create() (Lotto649AllSets, error) {
	tx := model.db.Create(model)
	return *model, tx.Error
}

// -
func (model *Lotto649AllSets) CreateInBatch(datas []Lotto649AllSets, batch_size int) error {
	tx := model.db.CreateInBatches(datas, batch_size)
	return tx.Error
}

func (model *Lotto649AllSets) Update(vals Lotto649AllSets) (Lotto649AllSets, error) {
	tx := model.db.Table("Lotto649AllSets").Updates(vals)
	return *model, tx.Error
}

func (model *Lotto649AllSets) Delete() error {
	tx := model.db.Where(1).Delete(model)
	return tx.Error
}

func (model *Lotto649AllSets) DeleteAll() error {
	tx := model.db.Where("1 = 1").Delete(&Lotto649AllSets{})
	return tx.Error
}

func (model *Lotto649AllSets) OrderByDESC(item_name string) *Lotto649AllSets {
	model.db = model.db.Order(clause.OrderByColumn{Column: clause.Column{Name: item_name}, Desc: true})
	return model
}
