package routes

import (
	"api/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	URI      string
	Metodo   string
	Callback gin.HandlerFunc
	Auth     bool
}

func Config(r *gin.Engine) {
	routes := userRoutes
	routes = append(routes, authRoutes...)
	for _, route := range routes {
		r.Handle(route.Metodo, route.URI, route.Callback)
	}
}

var authRoutes = []Route{
	{
		URI:      "/login",
		Metodo:   http.MethodPost,
		Callback: controllers.Login,
		Auth:     false,
	},
	// {
	// 	URI:      "/login",
	// 	Metodo:   http.MethodGet,
	// 	Callback: controllers.CreateUsers,
	// 	Auth:     false,
	// }, {
	// 	URI:      "/login",
	// 	Metodo:   http.MethodPost,
	// 	Callback: controllers.CreateUsers,
	// 	Auth:     false,
	// },
}
var userRoutes = []Route{
	{
		URI:      "/users",
		Metodo:   http.MethodPost,
		Callback: controllers.CreateUsers,
		Auth:     false,
	}, {
		URI:      "/users",
		Metodo:   http.MethodGet,
		Callback: controllers.GetUsers,
		Auth:     false,
	}, {
		URI:      "/users/:id",
		Metodo:   http.MethodDelete,
		Callback: controllers.DeleteUser,
		Auth:     false,
	}, {
		URI:      "/users/:id",
		Metodo:   http.MethodGet,
		Callback: controllers.ShowUser,
		Auth:     false,
	}, {
		URI:      "/users/:id",
		Metodo:   http.MethodPut,
		Callback: controllers.UpdateUser,
		Auth:     false,
	},
}
