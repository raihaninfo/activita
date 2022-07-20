package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("secret_key"))

func Auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "login-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, ok := session.Values["userId"]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		handlerFunc.ServeHTTP(w, r)
	}
}

func CheckLogin(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "login-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, ok := session.Values["userId"]
		if ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		handlerFunc.ServeHTTP(w, r)
	}
}
