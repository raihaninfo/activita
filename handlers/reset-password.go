package handlers

import (
	"fmt"
	"net/http"

	"github.com/raihaninfo/activita/helper"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

func RePassword(w http.ResponseWriter, r *http.Request) {
	views.ResetPassView.Template.Execute(w, nil)
}

func RePasswordAuth(w http.ResponseWriter, r *http.Request) {
	otp := r.FormValue("otp")
	newPassword := r.FormValue("newpass")
	confirmPassword := r.FormValue("conpass")

	if otp == ForgotOTP {
		if newPassword == confirmPassword {
			hash, err := helper.GenerateHash(newPassword, confirmPassword)
			if err != nil {
				fmt.Println(err)
			}

			err = models.UpdatePassword(UserEmail, hash)
			if err != nil {
				fmt.Println(err)
			}
			w.Header().Set("Location", "/login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			views.ResetPassView.Template.Execute(w, "Password Not Match")
		}

	} else {
		views.ResetPassView.Template.Execute(w, "OTP Not Match")
	}

}
