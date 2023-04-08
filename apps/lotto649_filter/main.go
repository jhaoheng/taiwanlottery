package main

import (
	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

var raw_results []model.Lottery

var op lotto649op.ILotto649OP

func main() {
	err := model.ConnSQLite("../../sql.db")
	if err != nil {
		panic(err)
	}

	raw_results, _ = model.NewLottery().FindAll()
	op = lotto649op.NewLotto649OP(raw_results)
	// module.GetAllPossiblility(op) // 取得所有組數

	/*
		[目前策略] 指定時間區間
		1. 排除掉連續 N 個數字, ex: 1,2,3,4
		2. 過濾掉 區間 的 中獎號碼
		3. 過濾掉 區間 的 歷史頭獎 取 5 個號碼, 再搭配 1 個號碼
		4. 過濾掉 區間 的 最後一次中獎號碼 7 個數字
		4. 取得 區間 後面的 N 次中獎號碼, 查看號碼在過濾後的資料中, 出現率 4 個以上機率多高
	*/

}
