package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbHost := "localhost:3306"
	dbPass := ""
	dbName := "rmediasolutions-website"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	return
}
