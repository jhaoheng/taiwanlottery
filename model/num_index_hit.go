package model

import (
	"fmt"

	"gorm.io/gorm"
)

type INumIndexHit interface {
	SetSID(sid int) *NumIndexHit
	SetNumIndexes(datas []NumIndex) *NumIndexHit
	Take() (NumIndexHit, error)
	FinaAll() ([]NumIndexHit, error)
	Create() (NumIndexHit, error)
	// 取得 <=sid 的指定數字總和
	Sum(sid, index int) (NumIndexHitSum, error)
}

type NumIndexHit struct {
	table_name string   `gorm:"-"`
	db         *gorm.DB `gorm:"-"`
	ID         int64
	SID        int `gorm:"column:sid"`
	NumIdx_1   int `gorm:"column:num_idx_1"` // 數字在 SID 中的數量排名
	NumIdx_2   int `gorm:"column:num_idx_2"`
	NumIdx_3   int `gorm:"column:num_idx_3"`
	NumIdx_4   int `gorm:"column:num_idx_4"`
	NumIdx_5   int `gorm:"column:num_idx_5"`
	NumIdx_6   int `gorm:"column:num_idx_6"`
	NumIdx_7   int `gorm:"column:num_idx_7"`
	NumIdx_8   int `gorm:"column:num_idx_8"`
	NumIdx_9   int `gorm:"column:num_idx_9"`
	NumIdx_10  int `gorm:"column:num_idx_10"`
	NumIdx_11  int `gorm:"column:num_idx_11"`
	NumIdx_12  int `gorm:"column:num_idx_12"`
	NumIdx_13  int `gorm:"column:num_idx_13"`
	NumIdx_14  int `gorm:"column:num_idx_14"`
	NumIdx_15  int `gorm:"column:num_idx_15"`
	NumIdx_16  int `gorm:"column:num_idx_16"`
	NumIdx_17  int `gorm:"column:num_idx_17"`
	NumIdx_18  int `gorm:"column:num_idx_18"`
	NumIdx_19  int `gorm:"column:num_idx_19"`
	NumIdx_20  int `gorm:"column:num_idx_20"`
	NumIdx_21  int `gorm:"column:num_idx_21"`
	NumIdx_22  int `gorm:"column:num_idx_22"`
	NumIdx_23  int `gorm:"column:num_idx_23"`
	NumIdx_24  int `gorm:"column:num_idx_24"`
	NumIdx_25  int `gorm:"column:num_idx_25"`
	NumIdx_26  int `gorm:"column:num_idx_26"`
	NumIdx_27  int `gorm:"column:num_idx_27"`
	NumIdx_28  int `gorm:"column:num_idx_28"`
	NumIdx_29  int `gorm:"column:num_idx_29"`
	NumIdx_30  int `gorm:"column:num_idx_30"`
	NumIdx_31  int `gorm:"column:num_idx_31"`
	NumIdx_32  int `gorm:"column:num_idx_32"`
	NumIdx_33  int `gorm:"column:num_idx_33"`
	NumIdx_34  int `gorm:"column:num_idx_34"`
	NumIdx_35  int `gorm:"column:num_idx_35"`
	NumIdx_36  int `gorm:"column:num_idx_36"`
	NumIdx_37  int `gorm:"column:num_idx_37"`
	NumIdx_38  int `gorm:"column:num_idx_38"`
	NumIdx_39  int `gorm:"column:num_idx_39"`
	NumIdx_40  int `gorm:"column:num_idx_40"`
	NumIdx_41  int `gorm:"column:num_idx_41"`
	NumIdx_42  int `gorm:"column:num_idx_42"`
	NumIdx_43  int `gorm:"column:num_idx_43"`
	NumIdx_44  int `gorm:"column:num_idx_44"`
	NumIdx_45  int `gorm:"column:num_idx_45"`
	NumIdx_46  int `gorm:"column:num_idx_46"`
	NumIdx_47  int `gorm:"column:num_idx_47"`
	NumIdx_48  int `gorm:"column:num_idx_48"`
	NumIdx_49  int `gorm:"column:num_idx_49"`
}

func NewNumIndexHit(table_name string) INumIndexHit {
	return &NumIndexHit{
		table_name: table_name,
		db:         db,
	}
}

func (model *NumIndexHit) SetSID(sid int) *NumIndexHit {
	model.SID = sid
	return model
}

type NumIndex struct {
	Index int
	Hit   int
}

func (model *NumIndexHit) SetNumIndexes(datas []NumIndex) *NumIndexHit {
	if len(datas) != 49 {
		err := fmt.Errorf("datas length wrong, %v", len(datas))
		panic(err)
	}

	//
	// data_map := map[int]bool{}
	for _, data := range datas {
		switch data.Index {
		case 1:
			model.NumIdx_1 = data.Hit
		case 2:
			model.NumIdx_2 = data.Hit
		case 3:
			model.NumIdx_3 = data.Hit
		case 4:
			model.NumIdx_4 = data.Hit
		case 5:
			model.NumIdx_5 = data.Hit
		case 6:
			model.NumIdx_6 = data.Hit
		case 7:
			model.NumIdx_7 = data.Hit
		case 8:
			model.NumIdx_8 = data.Hit
		case 9:
			model.NumIdx_9 = data.Hit
		case 10:
			model.NumIdx_10 = data.Hit
		case 11:
			model.NumIdx_11 = data.Hit
		case 12:
			model.NumIdx_12 = data.Hit
		case 13:
			model.NumIdx_13 = data.Hit
		case 14:
			model.NumIdx_14 = data.Hit
		case 15:
			model.NumIdx_15 = data.Hit
		case 16:
			model.NumIdx_16 = data.Hit
		case 17:
			model.NumIdx_17 = data.Hit
		case 18:
			model.NumIdx_18 = data.Hit
		case 19:
			model.NumIdx_19 = data.Hit
		case 20:
			model.NumIdx_20 = data.Hit
		case 21:
			model.NumIdx_21 = data.Hit
		case 22:
			model.NumIdx_22 = data.Hit
		case 23:
			model.NumIdx_23 = data.Hit
		case 24:
			model.NumIdx_24 = data.Hit
		case 25:
			model.NumIdx_25 = data.Hit
		case 26:
			model.NumIdx_26 = data.Hit
		case 27:
			model.NumIdx_27 = data.Hit
		case 28:
			model.NumIdx_28 = data.Hit
		case 29:
			model.NumIdx_29 = data.Hit
		case 30:
			model.NumIdx_30 = data.Hit
		case 31:
			model.NumIdx_31 = data.Hit
		case 32:
			model.NumIdx_32 = data.Hit
		case 33:
			model.NumIdx_33 = data.Hit
		case 34:
			model.NumIdx_34 = data.Hit
		case 35:
			model.NumIdx_35 = data.Hit
		case 36:
			model.NumIdx_36 = data.Hit
		case 37:
			model.NumIdx_37 = data.Hit
		case 38:
			model.NumIdx_38 = data.Hit
		case 39:
			model.NumIdx_39 = data.Hit
		case 40:
			model.NumIdx_40 = data.Hit
		case 41:
			model.NumIdx_41 = data.Hit
		case 42:
			model.NumIdx_42 = data.Hit
		case 43:
			model.NumIdx_43 = data.Hit
		case 44:
			model.NumIdx_44 = data.Hit
		case 45:
			model.NumIdx_45 = data.Hit
		case 46:
			model.NumIdx_46 = data.Hit
		case 47:
			model.NumIdx_47 = data.Hit
		case 48:
			model.NumIdx_48 = data.Hit
		case 49:
			model.NumIdx_49 = data.Hit
		}
	}
	return model
}

func (model *NumIndexHit) Take() (NumIndexHit, error) {
	output := NumIndexHit{}
	tx := model.db.Table(model.table_name).Where(model).Take(&output)
	return output, tx.Error
}

func (model *NumIndexHit) FinaAll() ([]NumIndexHit, error) {
	output := []NumIndexHit{}
	tx := model.db.Table(model.table_name).Where(model).Find(&output)
	return output, tx.Error
}

func (model *NumIndexHit) Create() (NumIndexHit, error) {
	tx := model.db.Table(model.table_name).Create(model)
	return *model, tx.Error
}

type NumIndexHitSum struct {
	Index int
	Total int
}

// 取得 <=sid 的指定數字總和
func (model *NumIndexHit) Sum(end_sid, index int) (NumIndexHitSum, error) {
	result := NumIndexHitSum{
		Index: index,
		Total: 0,
	}
	//
	w := fmt.Sprintf("`sid` <= %d AND `num_idx_%d`=1", end_sid, index)
	tx := model.db.Table(model.table_name).Select("sum(1) as total").Where(w).Find(&result)
	return result, tx.Error
}

func (model *NumIndexHit) ExportNumsToMap() map[int]int {
	output := map[int]int{
		1:  model.NumIdx_1,
		2:  model.NumIdx_2,
		3:  model.NumIdx_3,
		4:  model.NumIdx_4,
		5:  model.NumIdx_5,
		6:  model.NumIdx_6,
		7:  model.NumIdx_7,
		8:  model.NumIdx_8,
		9:  model.NumIdx_9,
		10: model.NumIdx_10,
		11: model.NumIdx_11,
		12: model.NumIdx_12,
		13: model.NumIdx_13,
		14: model.NumIdx_14,
		15: model.NumIdx_15,
		16: model.NumIdx_16,
		17: model.NumIdx_17,
		18: model.NumIdx_18,
		19: model.NumIdx_19,
		20: model.NumIdx_20,
		21: model.NumIdx_21,
		22: model.NumIdx_22,
		23: model.NumIdx_23,
		24: model.NumIdx_24,
		25: model.NumIdx_25,
		26: model.NumIdx_26,
		27: model.NumIdx_27,
		28: model.NumIdx_28,
		29: model.NumIdx_29,
		30: model.NumIdx_30,
		31: model.NumIdx_31,
		32: model.NumIdx_32,
		33: model.NumIdx_33,
		34: model.NumIdx_34,
		35: model.NumIdx_35,
		36: model.NumIdx_36,
		37: model.NumIdx_37,
		38: model.NumIdx_38,
		39: model.NumIdx_39,
		40: model.NumIdx_40,
		41: model.NumIdx_41,
		42: model.NumIdx_42,
		43: model.NumIdx_43,
		44: model.NumIdx_44,
		45: model.NumIdx_45,
		46: model.NumIdx_46,
		47: model.NumIdx_47,
		48: model.NumIdx_48,
		49: model.NumIdx_49,
	}
	return output
}
