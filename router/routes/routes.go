package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func Register(r *mux.Router) {
	// se tiver outras rotas (ex: booksRoutes), coloca aqui também
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
}
