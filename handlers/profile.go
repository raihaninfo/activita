package handlers

import (
	"net/http"

	"github.com/raihaninfo/activita/helper"
	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

type ProfileData struct {
	UserName     string
	Email        string
	SessionValue interface{}
	Error        string
}

func Profile(w http.ResponseWriter, r *http.Request) {
	sessionValue := middleware.SessionCheck(w, r)
	u, _ := models.UserById(sessionValue.(int))

	data := ProfileData{
		UserName:     u.Username,
		Email:        u.Email,
		SessionValue: sessionValue,
	}

	views.ProfileView.Template.Execute(w, data)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	session := middleware.SessionCheck(w, r)
	data := ProfileData{
		SessionValue: session,
	}
	views.DeleteAccountView.Template.Execute(w, data)
}

func DeleteAccountAuth(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("password")

	session := middleware.SessionCheck(w, r)
	data := ProfileData{
		Error:        "",
		SessionValue: session,
	}

	u, _ := models.UserById(session.(int))

	if u.Email == email {
		hash, err := helper.ComparePassword(u.Password, pass)
		if err != nil {
			data.Error = "Password is incorrect"
		}
		if hash {
			models.DeleteUserByEmail(email)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			data.Error = "Password is incorrect"
		}
	} else {
		data.Error = "Email not match"
	}

	views.DeleteAccountView.Template.Execute(w, data)
}
