package routes

import (
	"api/src/controllers"
	"api/src/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	URI      string
	Method   string
	Callback []gin.HandlerFunc
}

func Config(r *gin.Engine) {
	routes := [...][]Route{
		authRoutes,
		userRoutes,
		followRoutes,
	}
	for _, routeGroup := range routes {
		for _, route := range routeGroup {
			r.Handle(route.Method, route.URI, append([]gin.HandlerFunc{middlewares.Log}, route.Callback...)...)
		}
	}
}

var authRoutes = []Route{
	{
		URI:      "/login",
		Method:   http.MethodPost,
		Callback: []gin.HandlerFunc{controllers.Login},
	},
}
var userRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.CreateUsers},
	}, {
		URI:      "/users",
		Method:   http.MethodGet,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.GetUsers},
	}, {
		URI:      "/users/:id",
		Method:   http.MethodDelete,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.DeleteUser},
	}, {
		URI:      "/users/:id",
		Method:   http.MethodGet,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.ShowUser},
	}, {
		URI:      "/users/:id",
		Method:   http.MethodPut,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.UpdateUser},
	},
}

var followRoutes = []Route{
	{
		URI:      "/users/:id/follow",
		Method:   http.MethodPost,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.FollowUser},
	}, {
		URI:      "/users/:id/unfollow",
		Method:   http.MethodPost,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.UnFollowUser},
	}, {
		URI:      "/users/:id/followers",
		Method:   http.MethodGet,
		Callback: []gin.HandlerFunc{middlewares.Auth, controllers.GetFollowers},
	},
}
