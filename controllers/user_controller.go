package controller

import (
	"encoding/json"
	service "jokenpo-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	users, _ := service.FindAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UserController(router *mux.Router) {
	userRoute := router.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("", get).Methods(http.MethodGet)
}
