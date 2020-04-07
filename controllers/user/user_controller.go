package user

import (
	"encoding/json"
	user_repository "jokenpo-api/infrastructure/repository/user"
	user_service "jokenpo-api/services/user"
	"net/http"

	"github.com/gorilla/mux"
)

type userController struct {
	service user_service.UserService
}

type UserController interface {
	get(w http.ResponseWriter, r *http.Request)
}

func NewUserController() UserController {
	repository := user_repository.NewUserRepository()
	service := user_service.NewUserService(repository)

	return &userController{service}
}

func (uc userController) get(w http.ResponseWriter, r *http.Request) {
	users, _ := uc.service.FindAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func SetRouter(router *mux.Router) {
	uc := NewUserController()
	userRoute := router.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("", uc.get).Methods(http.MethodGet)
}
