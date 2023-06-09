package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lottery_lotto649(t *testing.T) {
	var b []byte

	// create
	fmt.Println("=== create ===")
	nums := Lotto649Nums{
		Num_1:      "10",
		Num_2:      "22",
		Num_3:      "26",
		Num_4:      "29",
		Num_5:      "45",
		Num_6:      "47",
		NumSpecial: "41",
	}
	b, _ = json.Marshal(nums)
	new_data, err := NewLottery().SetCategory(Lotto649).SetSerialID("112000038").SetBallNumbers(b).SetDate("112/03/31").Create()
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	b, _ = json.MarshalIndent(new_data, "", "	")
	fmt.Println(string(b))

	// take one
	fmt.Println("=== take ===")
	result, err := NewLottery().SetID(new_data.ID).Take()
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	b, _ = json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))

	// find all
	fmt.Println("=== find all ===")
	results, err := NewLottery().FindAll()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	b, _ = json.MarshalIndent(results, "", "	")
	fmt.Println(string(b))

	//
	fmt.Println("=== delete ===")
	if err := NewLottery().SetID(result.ID).Delete(); !assert.NoError(t, err) {
		t.Fatal()
	}
}

func Test_lottery_lotto649_orderby(t *testing.T) {
	//
	datas := []Lottery{
		0: {
			Category: Lotto649,
			SerialID: "112000001",
		},
		1: {
			Category: Lotto649,
			SerialID: "112000002",
		},
	}
	NewLottery().CreateInBatch(datas)
	//
	result, err := NewLottery().SetCategory(Lotto649).OrderByDESC("id").Take()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	b, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(b))
	//
	NewLottery().SetID(1).Delete()
	NewLottery().SetID(2).Delete()
}

func Test_lottery_lotto649_FindBetweenSID(t *testing.T) {
	results, _ := NewLottery().FindBetweenSID("112000041", "112000042")
	b, _ := json.MarshalIndent(results, "", "	")
	fmt.Println(string(b))
}
