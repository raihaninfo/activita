package handlers

import (
	"net/http"

	"github.com/raihaninfo/activita/views"
)

func Forgot(w http.ResponseWriter, r *http.Request) {
	views.ForgotView.Template.Execute(w, nil)
}
