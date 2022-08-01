package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func CreateLoginSession(w http.ResponseWriter, r *http.Request, thisUser int) {
	session, err := Store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30,
		HttpOnly: true,
	}
	session.Values["userId"] = thisUser
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SessionCheck(w http.ResponseWriter, r *http.Request) interface{} {
	session, err := Store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	_, ok := session.Values["userId"]
	if !ok {
		return nil
	}
	return session.Values["userId"]
}

// a function to delete the login session
func DeleteLoginSession(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	delete(session.Values, "userId")
	// session.Values["userId"] = nil
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}