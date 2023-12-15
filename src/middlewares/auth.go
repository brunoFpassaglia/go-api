package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	if error := auth.ValidateToken(c.Request); error != nil {
		responses.Error(c.Writer, http.StatusUnauthorized, error)
		c.Abort()
	}
}
