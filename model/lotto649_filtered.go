package model

import (
	"fmt"
	"sort"
	"time"

	"gorm.io/gorm"
)

/*
- 過濾後的可能數組
*/

type ILotto649Filtered interface {
	SetID(id int64) *Lotto649Filtered
	SetNums(nums []int) *Lotto649Filtered
	//
	Take() (Lotto649Filtered, error)
	FindAll() ([]Lotto649Filtered, error)
	FindNumsLike(texts []string) ([]Lotto649AllSets, error)
	Create() (Lotto649Filtered, error)
	CreateInBatch(datas []Lotto649Filtered, batch_size int) error
	Update(vals Lotto649Filtered) (Lotto649Filtered, error)
	Delete() error
	DeleteAll() error
}

type Lotto649Filtered struct {
	db        *gorm.DB  `gorm:"-"`
	ID        int64     `gorm:"primaryKey"`                             //
	Nums      string    `gorm:"index:idx_lotto649filtered_nums,unique"` // sort ascending and only 6 nums, ex: 1,2,3,4,5,6
	UpdatedAt time.Time //
	CreatedAt time.Time //
}

func (Lotto649Filtered) TableName() string {
	return "lotto649_filtered"
}

func NewLotto649Filtered() ILotto649Filtered {
	return &Lotto649Filtered{
		db: db,
	}
}

func (model *Lotto649Filtered) SetID(id int64) *Lotto649Filtered {
	model.ID = id
	return model
}

func (model *Lotto649Filtered) SetNums(nums []int) *Lotto649Filtered {
	if len(nums) != 6 {
		panic("nums length not 6")
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	model.Nums = fmt.Sprintf("%02d,%02d,%02d,%02d,%02d,%02d", nums[0], nums[1], nums[2], nums[3], nums[4], nums[5])
	return model
}

func (model *Lotto649Filtered) Take() (Lotto649Filtered, error) {
	output := Lotto649Filtered{}
	tx := model.db.Where(model).Take(&output)
	return output, tx.Error
}

func (model *Lotto649Filtered) FindAll() ([]Lotto649Filtered, error) {
	output := []Lotto649Filtered{}
	tx := model.db.Where(model).Find(&output)
	return output, tx.Error
}

// text, ex: %abc%
func (model *Lotto649Filtered) FindNumsLike(texts []string) ([]Lotto649AllSets, error) {
	output := []Lotto649AllSets{}
	tx := model.db.Where(model)
	for _, text := range texts {
		tx.Where("nums LIKE ?", text)
	}
	tx = tx.Find(&output)
	return output, tx.Error
}

func (model *Lotto649Filtered) Create() (Lotto649Filtered, error) {
	tx := model.db.Create(model)
	return *model, tx.Error
}

// -
func (model *Lotto649Filtered) CreateInBatch(datas []Lotto649Filtered, batch_size int) error {
	tx := model.db.CreateInBatches(datas, batch_size)
	return tx.Error
}

func (model *Lotto649Filtered) Update(vals Lotto649Filtered) (Lotto649Filtered, error) {
	tx := model.db.Table("Lotto649Filtered").Updates(vals)
	return *model, tx.Error
}

func (model *Lotto649Filtered) Delete() error {
	tx := model.db.Where(1).Delete(model)
	return tx.Error
}

func (model *Lotto649Filtered) DeleteAll() error {
	tx := model.db.Where("1 = 1").Delete(&Lotto649Filtered{})
	return tx.Error
}
