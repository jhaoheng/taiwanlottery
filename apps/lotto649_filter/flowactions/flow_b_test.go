package flowactions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/jhaoheng/taiwanlottery/lotto649op"
	"github.com/jhaoheng/taiwanlottery/model"
)

func Test_flowb_CalIndexHitSum(t *testing.T) {
	model.ConnMySQL()
	raw_results, _ := model.NewLottery().FindAll()
	op := lotto649op.NewLotto649OP(raw_results)
	//
	start_id := 1036
	flow_b := NewFlowB()
	rank_only_hit_indexes := flow_b.Run(start_id, 1, op)
	// fmt.Println(rank_only_hit_indexes)

	results := flow_b.CalIndexHitSum(rank_only_hit_indexes)
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}

func Test_flowb_import(t *testing.T) {
	model.ConnMySQL()

	filename := "./flow_b_import.csv"

	readFile, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for line_index, eachline := range fileTextLines {
		if line_index == 0 {
			continue
		}

		//
		num_indexes := []model.NumIndex{}
		sid := ""
		values := strings.Split(eachline, ",")
		for value_index, value := range values {
			if value_index == 0 {
				sid = value
				continue
			}
			num_indexes = append(num_indexes, model.NumIndex{
				Index: value_index,
				Hit: func() int {
					hit, _ := strconv.Atoi(value)
					return hit
				}(),
			})
		}
		// fmt.Println(sid, "==>", num_indexes)

		if _, err := model.NewNumIndexHit().SetSID(sid).SetNumIndexes(num_indexes).Create(); err != nil {
			panic(err)
		}
	}
}
