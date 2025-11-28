package routes

import (
	"Api-Aula1/controller"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:    "/users",
		Method: http.MethodPost, // CREATE
		Func:   controller.CreateUser,
	},
	{
		URI:    "/users",
		Method: http.MethodGet, // READ
		Func:   controller.FetchUser,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut, // UPDATE
		Func:   controller.UpdateUser,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete, // DELETE
		Func:   controller.DeleteUser,
	},
}
