package lotto649op

import (
	"fmt"
	"os"
	"testing"
)

func Test_NumsTrendingResult(t *testing.T) {
	filename, csv := NewLotto649OP(raw_results).ExportNumsTrending()

	//
	filepath := fmt.Sprintf("./nums_trending_output/%v", filename)
	os.WriteFile(filepath, []byte(csv), 0777)
}
