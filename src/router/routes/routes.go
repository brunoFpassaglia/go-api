package routes

import (
	"api/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	URI      string
	Metodo   string
	Callback func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Config(r *gin.Engine) {
	for _, route := range userRoutes {
		r.Handle(route.Metodo, route.URI, gin.WrapF(route.Callback))
	}
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
		URI:      "/users/{id}",
		Metodo:   http.MethodDelete,
		Callback: controllers.DeleteUser,
		Auth:     false,
	}, {
		URI:      "/users/{id}",
		Metodo:   http.MethodGet,
		Callback: controllers.ShowUser,
		Auth:     false,
	}, {
		URI:      "/users/{id}",
		Metodo:   http.MethodPut,
		Callback: controllers.UpdateUser,
		Auth:     false,
	},
}
