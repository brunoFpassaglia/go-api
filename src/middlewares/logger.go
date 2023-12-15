package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Log(c *gin.Context) {
	fmt.Println("\033[36m"+c.Request.Method, c.Request.RequestURI, c.Request.Host+"\033[0m")
	c.Next()
	fmt.Println("\033[36m" + "Request ended " + "\033[0m")
}
