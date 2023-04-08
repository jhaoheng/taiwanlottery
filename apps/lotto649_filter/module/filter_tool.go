package module

import (
	"fmt"
	"time"

	"github.com/jhaoheng/taiwanlottery/model"
)

/*
[目的]
- 帶入指定 datas, 在資料表中過濾掉 datas 的資料
*/

/*
	- 過濾 plan_a=> 取得 中獎號碼, 7 取 5 排列組合
	- 過濾{指定號碼} => datas := [][]string{0:{"1","2","3","4","5","6"}}
	- 過濾{最後一次的中獎號碼, 七個數字} => datas := M_Get_N([]int{},6), 取得排列組合
*/

type INewFilterTool interface {
	/*
		- 過濾{指定號碼} => datas := [][]string{0:{"1","2","3","4","5","6"}}
		- 過濾最後一次的中獎號碼, 七個數字 => datas := M_Get_N([]int{},6), 取得排列組合
	*/
	StartFilterByStr(datas [][]string)
	StartFilterByInt(datas [][]int)
}

type FilterTool struct {
}

func NewFilterTool() INewFilterTool {
	return &FilterTool{}
}

// -
func (filter *FilterTool) StartFilterByStr(datas [][]string) {
	start := time.Now()
	//
	del_count := 0
	for _, data := range datas {
		fmt.Printf("=> 查詢號碼: %v, ", data)
		if len(data) > 6 {
			panic("資料錯誤, 不得超過 6 個號碼")
		}
		//
		text := ""
		for _, n := range data {
			text = text + "%" + fmt.Sprintf("%02v", n) + "%"
		}
		finds, err := model.NewLotto649Filtered().FindNumsLike([]string{text})
		if err != nil {
			panic(err)
		}
		fmt.Printf("找到 %v\n", len(finds))
		if len(finds) != 0 {
			if err := model.NewLotto649Filtered().BatchDelete(finds); err != nil {
				panic(err)
			}
			del_count = del_count + len(finds)
		}
	}
	fmt.Printf("總共刪除: %v, 執行時間: %v\n", del_count, -time.Until(start))
}

// -
func (filter *FilterTool) StartFilterByInt(datas [][]int) {
	start := time.Now()
	//
	del_count := 0
	for _, data := range datas {
		fmt.Printf("=> 查詢號碼: %v, ", data)
		if len(data) > 6 {
			panic("資料錯誤, 不得超過 6 個號碼")
		}
		//
		text := ""
		for _, n := range data {
			text = text + "%" + fmt.Sprintf("%02d", n) + "%"
		}
		finds, err := model.NewLotto649Filtered().FindNumsLike([]string{text})
		if err != nil {
			panic(err)
		}
		fmt.Printf("找到 %v\n", len(finds))
		if len(finds) != 0 {
			if err := model.NewLotto649Filtered().BatchDelete(finds); err != nil {
				panic(err)
			}
			del_count = del_count + len(finds)
		}
	}
	fmt.Printf("總共刪除: %v, 執行時間: %v\n", del_count, -time.Until(start))
}
