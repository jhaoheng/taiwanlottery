package crawler

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Lotto649(t *testing.T) {
	result, err := NewLotto649(web_driver).SearchBySerialID("112000018")
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))
}
