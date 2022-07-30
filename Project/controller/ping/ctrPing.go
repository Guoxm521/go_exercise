package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping()gin.HandlerFunc  {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}