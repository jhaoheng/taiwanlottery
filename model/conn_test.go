package model

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	IsDebug = true
	ConnMySQL()

	m.Run()
	os.Exit(0)
}

func Test_Conn(t *testing.T) {
	IsDebug = true
	if err := ConnMySQL(); err != nil {
		t.Fatal()
	}
}
