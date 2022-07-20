package handlers

import (
	"fmt"
	"net/http"

	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

type UserData struct {
	SessionValue interface{}
	Verify       bool
}

func AddActivity(w http.ResponseWriter, r *http.Request) {
	sessionValue := middleware.SessionCheck(w, r)

	// userId := middleware.SessionCheck(w, r)

	user, err := models.UserById(sessionValue.(int))
	if err != nil {
		fmt.Println(err)
	}
	var status bool
	if user.Status ==0{
		status =false
	}else{
		status = true
	}
	var userData = UserData{
		SessionValue: sessionValue,
		Verify:       status,
	}

	views.AddView.Template.Execute(w, userData)

	fmt.Println(user.Status)
}

func AddActivityPost(w http.ResponseWriter, r *http.Request) {
	activity := r.FormValue("activity")
	description := r.FormValue("description")
	date := r.FormValue("date")
	time := r.FormValue("time")
	location := r.FormValue("location")
	category := r.FormValue("category")
	priority := r.FormValue("priority")
	// fmt.Println(activity, description, date, time, location, category, priority)

	userId := middleware.SessionCheck(w, r)

	models.NewActivity(userId, activity, description, date, time, location, category, priority)
	w.Header().Set("Location", "/")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
