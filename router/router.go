package router

import (
	"Api-Aula1/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	routes.Register(r)
	return r
}
