package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/R-Media-Solutions/rmediasolutions-website/config"
	"github.com/R-Media-Solutions/rmediasolutions-website/entities"
	"github.com/R-Media-Solutions/rmediasolutions-website/libraries"
	"github.com/R-Media-Solutions/rmediasolutions-website/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

var AdmUserModel = models.NewAdmUserModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"name": session.Values["name"],
			}

			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, data)
		}

	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// proses login
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		errorMessages := validation.Struct(UserInput)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)

		} else {

			var admuser entities.AdmUser
			AdmUserModel.Where(&admuser, "username", UserInput.Username)

			var message error
			if admuser.Username == "" {
				message = errors.New("Username atau Password salah!")
			} else {
				// pengecekan password
				errPassword := bcrypt.CompareHashAndPassword([]byte(admuser.Password), []byte(UserInput.Password))
				if errPassword != nil {
					message = errors.New("Username atau Password salah!")
				}
			}

			if message != nil {

				data := map[string]interface{}{
					"error": message,
				}

				temp, _ := template.ParseFiles("views/login.html")
				temp.Execute(w, data)
			} else {
				// set session
				session, _ := config.Store.Get(r, config.SESSION_ID)

				session.Values["loggedIn"] = true
				session.Values["email"] = admuser.Email
				session.Values["username"] = admuser.Username
				session.Values["name"] = admuser.Name

				session.Save(r, w)

				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}

	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// melakukan proses registrasi

		// mengambil inputan form
		r.ParseForm()

		admuser := entities.AdmUser{
			Name:      r.Form.Get("name"),
			Email:     r.Form.Get("email"),
			Username:  r.Form.Get("username"),
			Password:  r.Form.Get("password"),
			Cpassword: r.Form.Get("cpassword"),
		}

		errorMessages := validation.Struct(admuser)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       admuser,
			}

			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(admuser.Password), bcrypt.DefaultCost)
			admuser.Password = string(hashPassword)

			// insert ke database
			AdmUserModel.Create(admuser)

			data := map[string]interface{}{
				"pesan": "Registrasi berhasil",
			}
			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)
		}
	}

}
