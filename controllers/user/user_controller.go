package user

import (
	"encoding/json"
	"jokenpo-api/domain/model"
	"jokenpo-api/infrastructure/datastore"
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
	repository := user_repository.NewUserRepository(datastore.NewDB())
	service := user_service.NewUserService(repository)

	return &userController{service}
}

func (uc userController) get(w http.ResponseWriter, r *http.Request) {
	var u []*model.User
	users, _ := uc.service.FindAll(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func SetRouter(router *mux.Router) {
	uc := NewUserController()
	userRoute := router.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("", uc.get).Methods(http.MethodGet)
}
