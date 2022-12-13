package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	//Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}

// func DBConn() (db *sql.DB, err error) {
func DBConn() {
	//dbDriver := "mysql"
	dbUser := "root"
	dbHost := "localhost:3306"
	dbPass := ""
	dbName := "rmediasolutions-website"

	//db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	//log.Println("DB: ", db)

	Connect(dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName)
	Migrate()
	//return
}
