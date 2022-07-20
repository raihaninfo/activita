package handlers

import (
	"fmt"
	"net/http"

	"github.com/raihaninfo/activita/helper"
	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

type VerifyData struct {
	SessionValue interface{}
	Email        string
	Verify       bool
}

type VerifyDataPost struct {
	SessionValue interface{}
	Email        string
	Verify       bool
	Message      string
}

var OTPRandom string

func Verify(w http.ResponseWriter, r *http.Request) {
	sessionValue := middleware.SessionCheck(w, r)
	thisUser, _ := models.UserById(sessionValue.(int))
	var verify bool
	if thisUser.Status == 0 {
		verify = false
	} else {
		verify = true
	}

	if !verify {
		var err error
		OTPRandom, err = helper.GenerateOTP(6)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(OTPRandom)
		helper.SendMail(thisUser.Email, OTPRandom)
	}

	var verifyData = VerifyData{
		SessionValue: sessionValue,
		Email:        thisUser.Email,
		Verify:       verify,
	}

	views.VerifyView.Template.Execute(w, verifyData)
}

func VerifyAuth(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("verify-email")
	verifyCode := r.FormValue("verify-code")
	fmt.Println(email, verifyCode)
	sessionValue := middleware.SessionCheck(w, r)
	thisUser, _ := models.UserById(sessionValue.(int))
	var verify bool
	if thisUser.Status == 0 {
		verify = false
	} else {
		verify = true
	}
	var message string

	if verifyCode == OTPRandom {
		models.VerifyUser(thisUser.Id)
		w.Header().Set("Location", "/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		message = "Invalid Code"
	}

	var VerifyData = VerifyDataPost{
		SessionValue: sessionValue,
		Email:        thisUser.Email,
		Verify:       verify,
		Message:      message,
	}

	views.VerifyView.Template.Execute(w, VerifyData)
}

func VerifyReq(w http.ResponseWriter, r *http.Request) {
	sessionValue := middleware.SessionCheck(w, r)
	thisUser, _ := models.UserById(sessionValue.(int))
	var verify bool
	var VerifyData = VerifyDataPost{
		SessionValue: sessionValue,
		Email:        thisUser.Email,
		Verify:       verify,
		Message:      "",
	}

	if thisUser.Status == 0 {
		verify = false
	} else {
		verify = true
	}

	views.VerifyReqView.Template.Execute(w, VerifyData)

}
