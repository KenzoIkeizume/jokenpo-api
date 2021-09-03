package user

import (
	"encoding/json"
	"net/http"

	model "jokenpo-api/domain/model"
	datastore "jokenpo-api/infrastructure/datastore"
	user_repository "jokenpo-api/infrastructure/repository/user"
	user_input "jokenpo-api/input/controller/user"
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
	users, err := uc.service.FindAll(u)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"message": "error to list users"}`))
		return
	}

	users_response := []user_input.UserResponse{}
	for _, user_model := range users {
		users_response = append(users_response, user_input.UserResponse{
			ID:        user_model.ID,
			Name:      user_model.Name,
			Age:       user_model.Age,
			CreatedAt: user_model.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users_response)
}

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var u *user_input.UserRequest
	u = new(user_input.UserRequest)

	err := decoder.Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"message": "error to create user"}`))
		return
	}

	user_model := &model.User{
		Name: u.Name,
		Age:  u.Age,
	}

	user, err := uc.service.Create(user_model)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"message": "error when try to create user"}`))
		return
	}

	user_response := user_input.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user_response)
}

func SetRouter(router *mux.Router) {
	uc := NewUserController()
	userRoute := router.PathPrefix("/users").Subrouter()
	userRoute.HandleFunc("", uc.get).Methods(http.MethodGet)
	userRoute.HandleFunc("", uc.post).Methods(http.MethodPost)
}
