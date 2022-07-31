package ping

import (
	"example.com/m/v2/Project/model/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping()gin.HandlerFunc  {
	return func(c *gin.Context) {
		_account := db.NewAccount()
		fmt.Println("=======================================")
		fmt.Println(_account.GetByAccountID(2))
		fmt.Println(_account.Account)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}