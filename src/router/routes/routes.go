package routes

import (
	"api/src/controllers"
	"api/src/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	URI      string
	Metodo   string
	Callback []gin.HandlerFunc
}

func Config(r *gin.Engine) {
	routes := userRoutes
	routes = append(routes, authRoutes...)
	for _, route := range routes {
		r.Handle(route.Metodo, route.URI, append([]gin.HandlerFunc{middlewares.Log}, route.Callback...)...)
	}
}

var authRoutes = []Route{
	{
		URI:      "/login",
		Metodo:   http.MethodPost,
		Callback: []gin.HandlerFunc{controllers.Login},
	},
}
var userRoutes = []Route{
	{
		URI:      "/users",
		Metodo:   http.MethodPost,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.CreateUsers},
	}, {
		URI:      "/users",
		Metodo:   http.MethodGet,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.GetUsers},
	}, {
		URI:      "/users/:id",
		Metodo:   http.MethodDelete,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.DeleteUser},
	}, {
		URI:      "/users/:id",
		Metodo:   http.MethodGet,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.ShowUser},
	}, {
		URI:      "/users/:id",
		Metodo:   http.MethodPut,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.UpdateUser},
	},
}
