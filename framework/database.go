package frameworks

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const defaultTimeout = 2 * time.Second

type Database struct {
	*sql.DB
}

func LoadDatabase() (Database, error) {
	dsn := "root:example@tcp(localhost:3307)/movie_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("[DATABASE] error =>", err)
		return Database{}, err
	}

	ctx, cancel := defaultContext()
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("[DATABASE PING] error =>", err)
		return Database{}, err
	}

	return Database{db}, nil
}

func defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
