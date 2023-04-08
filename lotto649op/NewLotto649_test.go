package lotto649op

import (
	"fmt"
	"os"
	"testing"
	"time"

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

func Test_GetLotto649OPDatasWithFactor(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	start, _ := time.ParseInLocation("2006-01-02", "2014-01-01", loc)
	end, _ := time.ParseInLocation("2006-01-02", "2023-04-05", loc)

	results := NewLotto649OP(raw_results).GetLotto649OPDatasAndReplaceOne(start, end)
	// b, _ := json.MarshalIndent(results, "", "	")
	// fmt.Println("===>", string(b))
	fmt.Println("總共 =>", len(results))
}

/*
 */
func Test_GetLotto649OPDatas_CheckContinuousNums(t *testing.T) {
	// continuous_num := 3

}
