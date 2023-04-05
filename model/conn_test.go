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
