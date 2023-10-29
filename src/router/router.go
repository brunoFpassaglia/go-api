package router

import (
	"api/src/router/routes"

	"github.com/gin-gonic/gin"
)

func Gerar() *gin.Engine {
	router := gin.Default()
	routes.Config(router)
	return router
}
