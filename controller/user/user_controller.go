package user

import (
	"encoding/json"
	"net/http"

	model "jokenpo-api/domain/model"
	datastore "jokenpo-api/infrastructure/datastore"
	user_repository "jokenpo-api/infrastructure/repository/user"
	user_service "jokenpo-api/service/user"

	"github.com/gorilla/mux"
)

type userController struct {
	service user_service.UserService
}

type UserController interface {
	get(w http.ResponseWriter, r *http.Request)
	post(w http.ResponseWriter, r *http.Request)
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

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var u *model.User
	u = new(model.User)

	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	user, _ := uc.service.Create(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func SetRouter(router *mux.Router) {
	uc := NewUserController()
	userRoute := router.PathPrefix("/users").Subrouter()
	userRoute.HandleFunc("", uc.get).Methods(http.MethodGet)
	userRoute.HandleFunc("", uc.post).Methods(http.MethodPost)
}
