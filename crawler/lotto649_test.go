package crawler

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Lotto649(t *testing.T) {
	result := NewLotto649(web_driver).SearchBySerialID("112000018")
	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))
}
