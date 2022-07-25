package handlers

import (
	"net/http"

	"github.com/raihaninfo/activita/views"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	views.NotFoundView.Template.Execute(w, nil)
}
