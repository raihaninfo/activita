package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raihaninfo/activita/handlers"
	"github.com/raihaninfo/activita/middleware"
	"gorm.io/gorm"
)

func App(PORT int, db *gorm.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Auth(handlers.Home)).Methods(http.MethodGet)
	r.HandleFunc("/add", middleware.Auth(handlers.AddActivity)).Methods(http.MethodGet)
	r.HandleFunc("/add", middleware.Auth(handlers.AddActivityPost)).Methods(http.MethodPost)
	r.HandleFunc("/update/{id}", middleware.Auth(handlers.UpdateActivity)).Methods(http.MethodGet)
	r.HandleFunc("/update", middleware.Auth(handlers.UpdateActivityPost)).Methods(http.MethodPost)

	r.HandleFunc("/delete/{id}", middleware.Auth(handlers.DeleteActivity)).Methods(http.MethodGet)

	r.HandleFunc("/login", middleware.CheckLogin(handlers.Login)).Methods(http.MethodGet)
	r.HandleFunc("/login", handlers.LoginAuth).Methods(http.MethodPost)
	r.HandleFunc("/register", middleware.CheckLogin(handlers.Register)).Methods(http.MethodGet)
	r.HandleFunc("/register", handlers.RegisterAuth).Methods(http.MethodPost)
	r.HandleFunc("/forgot", handlers.Forgot).Methods(http.MethodGet)
	r.HandleFunc("/forgot", handlers.ForgotAuth).Methods(http.MethodPost)

	r.HandleFunc("/profile", middleware.Auth(handlers.Profile)).Methods(http.MethodGet)
	r.HandleFunc("/verify", handlers.Verify).Methods(http.MethodGet)
	r.HandleFunc("/verify", handlers.VerifyAuth).Methods(http.MethodPost)
	r.HandleFunc("/verify/req", handlers.VerifyReq).Methods(http.MethodGet)
	r.HandleFunc("/logout", middleware.Auth(handlers.Logout)).Methods(http.MethodGet)
	r.HandleFunc("/repass", handlers.RePassword).Methods(http.MethodGet)
	r.HandleFunc("/repass", handlers.RePasswordAuth).Methods(http.MethodPost)
	r.HandleFunc("/profile/delete", middleware.Auth(handlers.DeleteAccount)).Methods(http.MethodGet)
	r.HandleFunc("/profile/delete", middleware.Auth(handlers.DeleteAccountAuth)).Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	log.Printf("Listening Port: %v", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), r)
}
