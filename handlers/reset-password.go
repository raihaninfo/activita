package handlers

import (
	"net/http"

	"github.com/raihaninfo/activita/views"
)

func RePassword(w http.ResponseWriter, r *http.Request) {
	views.ResetPassView.Template.Execute(w, nil)
}
