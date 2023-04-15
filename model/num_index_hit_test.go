package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumIndexHit_FindAll(t *testing.T) {
	results, err := NewNumIndexHit().FinaAll()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Printf("%+v\n", results[0])
}

func Test_NumIndexHit_Create(t *testing.T) {
	num_indexes := []NumIndex{}
	for i := 1; i <= 49; i++ {
		is_hit := 0
		if i%2 == 0 {
			is_hit = 1
		}

		num_indexes = append(num_indexes,
			NumIndex{
				Index: i,
				Hit:   is_hit,
			})
	}
	item, err := NewNumIndexHit().SetSID(12345).SetNumIndexes(num_indexes).Create()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Printf("%+v\n", item)
}

func Test_NumIndexHit_Sum_1(t *testing.T) {
	result, _ := NewNumIndexHit().Sum(112000040, 2)
	fmt.Println(result)
}

func Test_NumIndexHit_SumTreanding(t *testing.T) {

	// sums := []NumIndexHitSum{}
	for i := 1; i <= 49; i++ {
		sum, _ := NewNumIndexHit().Sum(112000040, i)
		fmt.Println(sum)
	}
}
