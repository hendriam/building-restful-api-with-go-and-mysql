package test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnectionDatabase(t *testing.T) {
	dsn := "root:example@tcp(localhost:3307)/movie_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
