package handlers

import (
	"fmt"
	"net/http"

	"github.com/raihaninfo/activita/helper"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

func Forgot(w http.ResponseWriter, r *http.Request) {
	views.ForgotView.Template.Execute(w, nil)
}

type ForgotData struct {
	User bool
}

var ForgotOTP string
var UserEmail string

func ForgotAuth(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	user, err := models.UserByEmail(email)
	if err != nil {
		fmt.Println(err)
	}
	UserEmail = user.Email
	data := ForgotData{}
	if user.Email == "" {
		data.User = false
	} else {
		data.User = true
		http.Redirect(w, r, "/repass", http.StatusSeeOther)
		ForgotOTP, _ = helper.GenerateOTP(6)

		message := "<body><h3>Dear User!</h3> <br>your Forgot password otp is <h2 style=\"text-align:center;\"><span style=\"font-size:40px;border:2px solid black;padding:10px\">%v</span></h2> \n</body>"

		helper.SendMail(user.Email, ForgotOTP, "Subject: Activita Forgot Password\n", message)
	}
	views.ForgotView.Template.Execute(w, data)
}
