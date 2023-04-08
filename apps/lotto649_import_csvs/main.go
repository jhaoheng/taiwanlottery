package main

import (
	"encoding/json"
	"fmt"

	"github.com/jhaoheng/taiwanlottery/model"

	"gorm.io/gorm"
)

func main() {
	ImportLotto649()
	// ImportSuperLotto638()
}

/*
- 判斷資料是否存在
- 寫入新資料
*/
func WriteToDB(category model.LotteryCategory, lotterys []model.Lottery) {
	model.ConnMySQL()
	db_write_success := 0
	for _, lottery := range lotterys {
		obj, err := model.NewLottery().SetCategory(category).SetSerialID(lottery.SerialID).Take()
		if err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		if len(obj.SerialID) != 0 {
			continue
		}

		if _, err := model.NewLotteryWith(lottery).Create(); err != nil {
			b, _ := json.MarshalIndent(lottery, "", "	")
			fmt.Println(string(b))
			panic(err)
		}
		db_write_success++
	}

	fmt.Printf("寫入資料庫有 %v 筆\n", db_write_success)
}
