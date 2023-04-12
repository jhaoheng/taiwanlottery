package model

import (
	"gorm.io/gorm"
)

/*
- 過濾後的可能數組
*/

/*
- 複製資料: INSERT INTO lotto649_filtered SELECT * FROM lotto649_all_sets;
*/

type ILotto649Filtered interface {
	SetID(id int64) *Lotto649Filtered
	SetNums(nums string) *Lotto649Filtered
	//
	Take() (Lotto649Filtered, error)
	FindAll() ([]Lotto649Filtered, error)
	FindNumsLike(texts []string) ([]Lotto649Filtered, error)
	Create() (Lotto649Filtered, error)
	CreateInBatch(datas []Lotto649Filtered, batch_size int) error
	Update(vals Lotto649Filtered) (Lotto649Filtered, error)
	Delete() error
	DeleteAll() error
	BatchDelete(objs []Lotto649Filtered) error
}

type Lotto649Filtered struct {
	db        *gorm.DB `gorm:"-"`
	TableName string   `gorm:"-"`
	ID        int64    `gorm:"primaryKey;type:int(11);autoIncrement;not null"` //
	Nums      string   `gorm:"type:varchar(20);uniqueIndex;not null;"`         // sort ascending and only 6 nums, ex: 1,2,3,4,5,6
}

func NewLotto649Filtered(table_name string) ILotto649Filtered {
	// 檢查 table 是否存在, 不存在則新增
	if !db.Migrator().HasTable(table_name) {
		migrator := db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator()
		if err := migrator.CreateTable(&Lotto649Filtered{}); err != nil {
			panic(err)
		}
		if err := migrator.RenameTable("lotto649_filtereds", table_name); err != nil {
			panic(err)
		}
	} else {
		panic("table 已經存在, 請確認是否要移除")
	}

	return &Lotto649Filtered{
		TableName: table_name,
		db:        db,
	}
}

func (model *Lotto649Filtered) SetID(id int64) *Lotto649Filtered {
	model.ID = id
	return model
}

func (model *Lotto649Filtered) SetNums(nums string) *Lotto649Filtered {
	model.Nums = nums
	return model
}

func (model *Lotto649Filtered) Take() (Lotto649Filtered, error) {
	output := Lotto649Filtered{}
	tx := model.db.Table(model.TableName).Where(model).Take(&output)
	return output, tx.Error
}

func (model *Lotto649Filtered) FindAll() ([]Lotto649Filtered, error) {
	output := []Lotto649Filtered{}
	tx := model.db.Table(model.TableName).Where(model).Find(&output)
	return output, tx.Error
}

// text, ex: %abc%
func (model *Lotto649Filtered) FindNumsLike(texts []string) ([]Lotto649Filtered, error) {
	output := []Lotto649Filtered{}
	tx := model.db.Table(model.TableName).Where(model)

	for _, text := range texts {
		tx = tx.Or("nums LIKE ?", text)
	}
	tx = tx.Find(&output)
	return output, tx.Error
}

func (model *Lotto649Filtered) Create() (Lotto649Filtered, error) {
	tx := model.db.Table(model.TableName).Create(model)
	return *model, tx.Error
}

// -
func (model *Lotto649Filtered) CreateInBatch(datas []Lotto649Filtered, batch_size int) error {
	tx := model.db.Table(model.TableName).CreateInBatches(datas, batch_size)
	return tx.Error
}

func (model *Lotto649Filtered) Update(vals Lotto649Filtered) (Lotto649Filtered, error) {
	tx := model.db.Table(model.TableName).Updates(vals)
	return *model, tx.Error
}

func (model *Lotto649Filtered) Delete() error {
	tx := model.db.Table(model.TableName).Where(model).Delete(&Lotto649Filtered{})
	return tx.Error
}

func (model *Lotto649Filtered) DeleteAll() error {
	tx := model.db.Table(model.TableName).Where("1=1").Delete(&Lotto649Filtered{})
	return tx.Error
}

func (model *Lotto649Filtered) BatchDelete(objs []Lotto649Filtered) error {
	tx := model.db.Table(model.TableName).Delete(objs)
	return tx.Error
}
