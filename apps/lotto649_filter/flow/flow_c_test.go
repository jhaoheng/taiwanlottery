package flow

import (
	"strconv"
	"testing"

	"github.com/jhaoheng/taiwanlottery/apps/lotto649_filter/module"
	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_flowc_RunPlansAndGetReports(t *testing.T) {
	model.ConnMySQL()
	_, all_sets_map := module.NewAllSets().Get()

	inferential_sid := 112000041
	flow_c := NewFlowC(inferential_sid, all_sets_map)
	csv_datas := flow_c.RunPlansAndGetReports()
	flow_c.SaveReports(csv_datas, "flow_c_test_"+strconv.Itoa(inferential_sid))
}
