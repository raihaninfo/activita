package handlers

import (
	"net/http"

	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/views"
)

type ActivitiesData struct {
	SessionValue interface{}
	Activities   []models.Activity
}

func Home(w http.ResponseWriter, r *http.Request) {
	sessionValue := middleware.SessionCheck(w, r)
	id, _ := models.AllActivityById(sessionValue.(int))

	var homeData = ActivitiesData{
		SessionValue: sessionValue,
		Activities:  id,
	}

	views.HomeView.Template.Execute(w, homeData)
	// fmt.Println(sessionValue, jj)
}
