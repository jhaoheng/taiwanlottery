package module

import (
	"fmt"
	"time"

	"github.com/jhaoheng/taiwanlottery/model"
)

/*
- 排除 => 歷史頭獎 取 5 個號碼, 再搭配 1 個號碼
- sets 數字, ascending
- remove_sets, ascending
*/

type INewFilterPlanB interface {
	// datas := NewLotto649OP(raw_results).GetLotto649OPDatasAndReplaceOne(start, end)
	StartFilter(datas [][]string)
}

type FilterPlanB struct {
}

func NewFilterPlanB() INewFilterPlanB {
	return &FilterPlanB{}
}

/*
- datas: 從 NewLotto649OP 中, 取得計算得到的數字
*/
func (filter *FilterPlanB) StartFilter(datas [][]string) {
	start := time.Now()
	//
	del_count := 0
	for _, data := range datas {
		fmt.Printf("=> 查詢號碼: %v, ", data)
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
