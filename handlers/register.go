package handlers

import (
	"log"
	"net/http"

	"github.com/raihaninfo/activita/helper"
	"github.com/raihaninfo/activita/message"
	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

func Register(w http.ResponseWriter, r *http.Request) {
	views.RegisterView.Template.Execute(w, nil)
}

func RegisterAuth(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	hashPass, err := helper.GenerateHash(password, confirmPassword)
	if err != nil {
		log.Fatal(err)
	}

	msg := message.RegisterMessage(username, email, password, confirmPassword)
	if msg.Username == "" && msg.Email == "" && msg.Password == "" {
		user, err := models.NewUser(username, email, hashPass)
		if err != nil {
			panic(err)
		}
		thisUser, _ := models.UserById(user.Id)
		middleware.CreateLoginSession(w, r, thisUser.Id)
		w.Header().Set("Location", "/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	views.RegisterView.Template.Execute(w, msg)
}
