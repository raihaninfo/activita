package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

type UserData struct {
	SessionValue interface{}
	Verify       bool
}
type UpdateActivityData struct {
	SessionValue interface{}
	Verify       bool
	Activity     models.Activity
}

func AddActivity(w http.ResponseWriter, r *http.Request) {
	sessionValue := middleware.SessionCheck(w, r)

	// userId := middleware.SessionCheck(w, r)

	user, err := models.UserById(sessionValue.(int))
	if err != nil {
		fmt.Println(err)
	}
	var status bool
	if user.Status == 0 {
		status = false
	} else {
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

// update activity
func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// fmt.Println(id)
	intId, _ := strconv.Atoi(id)
	activity, err := models.ActivityById(intId)
	if err != nil {
		fmt.Println(err)
	}

	sessionValue := middleware.SessionCheck(w, r)
	// userId := middleware.SessionCheck(w, r)
	var userData = UpdateActivityData{
		SessionValue: sessionValue,
		Verify:       true,
		Activity:     *activity,
	}

	views.UpdateView.Template.Execute(w, userData)
}

func UpdateActivityPost(w http.ResponseWriter, r *http.Request) {
	// update activity
	id := r.FormValue("id")
	Activity := r.FormValue("activity")
	Description := r.FormValue("description")
	Date := r.FormValue("date")
	Time := r.FormValue("time")
	Location := r.FormValue("location")
	Category := r.FormValue("category")
	Priority := r.FormValue("priority")

	intId, _ := strconv.Atoi(id)

	models.UpdateActivity(intId, Activity, Description, Date, Time, Location, Category, Priority)
	w.Header().Set("Location", "/")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
