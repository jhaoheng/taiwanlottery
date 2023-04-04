package model

import "testing"

func Test_Conn(t *testing.T) {
	path := "./testdb/sql.db"
	ConnSQLite(path)
}
