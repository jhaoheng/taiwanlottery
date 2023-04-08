package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
