package crawler

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Superlotto638(t *testing.T) {
	result, err := NewSuperlotto638(web_driver).SearchBySerialID("112000026")
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))
}
