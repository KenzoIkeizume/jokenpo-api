package controller

import (
	"net/http"

	config "jokenpo-api/config"
	user_controller "jokenpo-api/controller/user"

	"github.com/golang/glog"
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

	glog.Info("Server started on port %s", config.C.Server.Address)

	err := http.ListenAndServe(":"+config.C.Server.Address, router)
	if err != nil {
		glog.Fatal("Server cannot listen: %+v", err)
	}
}
