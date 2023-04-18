package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/stat"
)

func Test_rank_create(t *testing.T) {
	datas := []RankData{
		0: {Num: 1, Count: 1},
		1: {Num: 2, Count: 0},
	}
	b, _ := json.Marshal(datas)
	_, err := NewRank().SetSID(123).SetData(b).Create()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
}

func Test_rank_Take(t *testing.T) {
	item, _ := NewRank().SetSID(123).Take()
	fmt.Println(item)
}

func Test_rank_Variance(t *testing.T) {
	item, _ := NewRank().SetSID(106000024).Take()
	datas := []RankData{}
	json.Unmarshal([]byte(item.Data), &datas)

	// numbers := []float64{}
	weights := []float64{}
	for _, data := range datas {
		// numbers = append(numbers, float64(index))
		weights = append(weights, float64(data.Count))
	}
	mean := stat.Mean(weights, nil)
	variance := stat.Variance(weights, nil)
	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Variance: %.2f\n", variance)
}
