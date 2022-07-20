package handlers

import (
	"log"
	"net/http"

	"github.com/raihaninfo/activita/helper"
	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

func Login(w http.ResponseWriter, r *http.Request) {
	views.LoginView.Template.Execute(w, nil)
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := models.UserByUsername(username)
	if err != nil {
		log.Println(err)
	}
	var msg string
	if user.Username == username {
		ok, err := helper.ComparePassword(user.Password, password)
		if err != nil {
			msg = "Password doesn't match"
			log.Println(err)
		}
		if ok {
			middleware.CreateLoginSession(w, r, user.Id)
			w.Header().Set("Location", "/")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		msg = "Username doesn't exist"
	}
	views.LoginView.Template.Execute(w, msg)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	// logout
	middleware.DeleteLoginSession(w, r)
	w.Header().Set("Location", "/")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
