package controller

import (
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
	UserController(router)
	router.HandleFunc("", notFound)

	log.Fatal(http.ListenAndServe(":8080", router))
}