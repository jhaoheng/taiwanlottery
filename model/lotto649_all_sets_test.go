package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Lotto649AllSets_serial(t *testing.T) {
	fmt.Printf("\n\n=== create item ===")
	new_item, err := NewLotto649AllSets().SetNums([]int{1, 2, 3, 4, 6, 5}).Create()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Printf("===> create success\n\n")

	fmt.Printf("=== get item ===")
	get_item, err := NewLotto649AllSets().SetID(new_item.ID).Take()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Printf("%v\n\n", func() string {
		b, _ := json.MarshalIndent(get_item, "", "	")
		return string(b)
	}())

	fmt.Printf("=== find all item ===")
	find_items, err := NewLotto649AllSets().FindAll()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Printf("%v\n\n", func() string {
		b, _ := json.MarshalIndent(find_items, "", "	")
		return string(b)
	}())

	fmt.Printf("=== del item ===")
	err = NewLotto649AllSets().DeleteAll()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Printf("===> del all success\n\n")
}

// func Test_Lotto649AllSets_FindNumsLike(t *testing.T) {
// 	results, err := NewLotto649AllSets().FindNumsLike([]string{"%%02%%03%%08%%12%%"})
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}

// 	b, _ := json.MarshalIndent(results, "", "	")
// 	fmt.Println(string(b))
// }
