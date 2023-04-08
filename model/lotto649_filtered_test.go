package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Lotto649Filtered_FindNumsLike(t *testing.T) {
	texts := []string{
		"%%01%%02%%03%%04%%15%%35%%",
		"%%01%%02%%03%%04%%15%%36%%",
	}
	results, err := NewLotto649Filtered().FindNumsLike(texts)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	b, _ := json.MarshalIndent(results, "", "	")
	fmt.Println(string(b))
}

func Test_Lotto649Filtered_BatchDelete(t *testing.T) {
	objs := []Lotto649Filtered{
		0: {ID: 3},
		1: {ID: 4},
	}
	err := NewLotto649Filtered().BatchDelete(objs)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
}
