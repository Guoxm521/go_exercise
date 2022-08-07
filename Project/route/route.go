package route

import (
	"example.com/m/v2/controller/github"
	"example.com/m/v2/controller/ping"
	"example.com/m/v2/database"
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
	router.GET("/github/trending", github.GithubTrendingList())
}
