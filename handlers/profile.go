package handlers

import (
	"net/http"

	"github.com/raihaninfo/activita/middleware"
	"github.com/raihaninfo/activita/views"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	sessionValue := middleware.SessionCheck(w, r)
	views.ProfileView.Template.Execute(w, sessionValue)
}
