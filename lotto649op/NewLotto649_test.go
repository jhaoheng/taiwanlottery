package lotto649op

import (
	"os"
	"testing"

	"github.com/jhaoheng/taiwanlottery/model"
)

var raw_results []model.Lottery

func TestMain(m *testing.M) {
	err := model.ConnSQLite("../sql.db")
	if err != nil {
		panic(err)
	}

	raw_results, _ = model.NewLottery().FindAll()

	m.Run()
	os.Exit(0)
}
