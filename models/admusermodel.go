package models

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

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
	}

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

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
