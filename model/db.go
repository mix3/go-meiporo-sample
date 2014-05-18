package model

import (
	"fmt"
	"os"

	"github.com/naoina/genmai"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var db *genmai.DB

// for test
var initDB = func() *genmai.DB {
	driver := fmt.Sprintf("%s", os.Getenv("DRIVER"))
	dsn := fmt.Sprintf("%s", os.Getenv("DSN"))
	var d genmai.Dialect
	switch driver {
	case "mysql":
		d = &genmai.MySQLDialect{}
	case "postgres":
		d = &genmai.PostgresDialect{}
	case "sqlite3":
		d = &genmai.SQLite3Dialect{}
	default:
		panic(fmt.Errorf("genmai: unsupported driver type: %v", driver))
	}
	var err error
	db, err = genmai.New(d, dsn)
	if err != nil {
		panic(err)
	}
	db.CreateTableIfNotExists(&Todo{})
	return db
}

func GetDB() *genmai.DB {
	if db != nil {
		return db
	}
	db = initDB()
	return db
}
