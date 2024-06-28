package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"rutube/internal/auth"
	"rutube/internal/handlers"
)

func RegisterRoutes(birthdayHandler *handlers.BirthdayHandler) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/birthdays", auth.Middleware(http.HandlerFunc(birthdayHandler.GetTodaysBirthdays))).Methods("GET")
	router.HandleFunc("/register", birthdayHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", birthdayHandler.Login).Methods("POST")
	router.HandleFunc("/subscribe", birthdayHandler.Subscribe).Methods("POST")
	router.HandleFunc("/unsubscribe", birthdayHandler.Unsubscribe).Methods("POST")
	return router
}
