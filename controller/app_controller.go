package controller

import (
	"log"
	"net/http"

	config "jokenpo-api/config"
	user_controller "jokenpo-api/controller/user"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func AppController() {
	router := mux.NewRouter()
	user_controller.SetRouter(router)
	router.HandleFunc("", notFound)
	println("address: ", config.C.Server.Address)
	log.Fatal(http.ListenAndServe(":"+config.C.Server.Address, router))
}
