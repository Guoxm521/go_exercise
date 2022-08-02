package route

import (
	"example.com/m/v2/Project/controller/github"
	"example.com/m/v2/Project/controller/ping"
	"example.com/m/v2/Project/database"
	"example.com/m/v2/Project/model"
	"example.com/m/v2/Project/model/db"
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
