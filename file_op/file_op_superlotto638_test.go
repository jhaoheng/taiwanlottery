package fileop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileop_ParsedSuperlotto638(t *testing.T) {
	filepath := "../docs/2014/威力彩_2014.csv"
	file_op, err := NewFileOP().Read(filepath)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	results, err := file_op.ParsedSuperlotto638(",")
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}
