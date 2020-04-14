package controller

import (
	"jokenpo-api/config"
	userController "jokenpo-api/controllers/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func AppController() {
	router := mux.NewRouter()
	userController.SetRouter(router)
	router.HandleFunc("", notFound)
	println("address: ", config.C.Server.Address)
	log.Fatal(http.ListenAndServe(":"+config.C.Server.Address, router))
}
