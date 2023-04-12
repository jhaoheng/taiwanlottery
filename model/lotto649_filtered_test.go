package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Lotto649Filtered_NewLotto649Filtered(t *testing.T) {
	table_name := "ta123"
	db.Migrator().DropTable(table_name)
	filtered := NewLotto649Filtered(table_name)
	//
	fmt.Println("=== create")
	new_item, err := filtered.SetNums("1,2,3,4,5,6").Create()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Println(func() string {
		b, _ := json.MarshalIndent(new_item, "", "	")
		return string(b)
	}())
	//
	fmt.Println("=== take one")
	take_one, err := filtered.SetID(new_item.ID).Take()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	fmt.Println(func() string {
		b, _ := json.MarshalIndent(take_one, "", "	")
		return string(b)
	}())
	//
	fmt.Println("=== delete one")
	err = filtered.SetID(new_item.ID).Delete()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
}

func Test_Lotto649Filtered_FindNumsLike(t *testing.T) {
	texts := []string{
		"%%01%%02%%03%%04%%15%%35%%",
		"%%01%%02%%03%%04%%15%%36%%",
	}
	results, err := NewLotto649Filtered("").FindNumsLike(texts)
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
	err := NewLotto649Filtered("").BatchDelete(objs)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
}
