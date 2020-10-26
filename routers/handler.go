package routers

import (
	"github.com/gorilla/mux"
)

func Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/test", GetTest).Methods("GET")

	return r
}
