package models

import (
	"database/sql"
	"log"

	"github.com/R-Media-Solutions/rmediasolutions-website/config"
	"github.com/R-Media-Solutions/rmediasolutions-website/entities"
)

type AdmUserModel struct {
	db *sql.DB
}

func NewAdmUserModel() *AdmUserModel {
	conn, err := config.DBConn()

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")

	return &AdmUserModel{
		db: conn,
	}
}

func (u AdmUserModel) Where(admuser *entities.AdmUser, fieldName, fieldValue string) error {

	row, err := u.db.Query("SELECT id, name, email, username, password FROM adm_users WHERE "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&admuser.Id, &admuser.Name, &admuser.Email, &admuser.Username, &admuser.Password)
	}

	return nil
}

func (u AdmUserModel) Create(admuser entities.AdmUser) (int64, error) {

	result, err := u.db.Exec("INSERT INTO adm_users (name, email, username, password) values(?,?,?,?)",
		admuser.Name, admuser.Email, admuser.Username, admuser.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil
}
