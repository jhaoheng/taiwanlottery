package flow

import (
	"strconv"
	"testing"

	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_flowc_RunPlansAndGetReports(t *testing.T) {
	model.ConnMySQL()

	inferential_sid := 112000042
	flow_c := NewFlowC(inferential_sid)
	csv_datas := flow_c.RunPlansAndGetReports()
	flow_c.SaveReports(csv_datas, "flow_c_test_"+strconv.Itoa(inferential_sid))
}
