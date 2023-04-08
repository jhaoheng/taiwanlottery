package model

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	path := "./testdb/sql.db"
	IsDebug = true
	ConnSQLite(path)

	m.Run()
	os.Exit(0)
}

func Test_Conn(t *testing.T) {
	path := "./testdb/sql.db"
	IsDebug = true
	if err := ConnSQLite(path); err != nil {
		t.Fatal()
	}
}
