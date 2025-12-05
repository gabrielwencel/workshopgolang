package routes

import (
	"Api-Aula1/controller"
	"net/http"
)

var usersRoutes = []Routes{
	// Criar usuário
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
	},

	// Buscar todos os usuários
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controller.FetchUser,
	},

	// Buscar um usuário por ID
	{
		URI:    "/users/{userID}",
		Method: http.MethodGet,
		Func:   controller.FetchUser,
	},

	// Atualizar usuário
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
	},

	// Deletar usuário
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
	},

	// 🔥 NOVA ROTA: LOGIN
	{
		URI:    "/login",
		Method: http.MethodPost,
		Func:   controller.Login,
	},
}
