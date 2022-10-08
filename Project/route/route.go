package route

import (
	"example.com/m/v2/controller/account"
	"example.com/m/v2/controller/github"
	"example.com/m/v2/controller/ping"
	"example.com/m/v2/database"
	"example.com/m/v2/middleware"
	"example.com/m/v2/model"
	"example.com/m/v2/model/db"
	"github.com/gin-gonic/gin"
)

func config() {
	db.SetORM(database.OrmEngine)
	model.GoFunc(model.NewSync().Cron())
}

func RouterApi(router *gin.Engine) {
	config()
	router.GET("/ping", ping.Ping())
	router.POST("/account/login", account.AccountLogin())
	router.GET("/github/trending", github.GithubTrendingList())
	jwt := router.Group("/")
	jwt.Use(middleware.JWT())
	{
		jwt.GET("/ping1", ping.Ping())
		jwt.POST("/account/add", account.AccountAdd())
	}

}
