package fileop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileop_ParsedLotto649CSV(t *testing.T) {
	filepath := "../taiwan_lotto_csvs/2014/大樂透_2014.csv"
	file_op, err := NewFileOP().Read(filepath)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	results, err := file_op.ParsedLotto649CSV(",")
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}
