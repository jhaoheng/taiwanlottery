package crawler

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Superlotto638(t *testing.T) {
	result := NewSuperlotto638(web_driver).SearchBySerialID("112000026")
	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))
}
