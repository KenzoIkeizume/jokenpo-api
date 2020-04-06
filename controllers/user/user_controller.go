package user

import (
	"encoding/json"
	userService "jokenpo-api/services/user"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	users, _ := userService.FindAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func SetRouter(router *mux.Router) {
	userRoute := router.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("", get).Methods(http.MethodGet)
}
