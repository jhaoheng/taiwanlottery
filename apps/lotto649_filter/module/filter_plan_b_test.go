package module

import (
	"testing"

	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_FilterPlanB(t *testing.T) {
	model.IsDebug = true
	model.ConnSQLite("../../../sql.db")
	datas := [][]string{
		0: {"01", "02", "03", "04", "05", "06"},
	}
	NewFilterPlanB().StartFilter(datas)
}
