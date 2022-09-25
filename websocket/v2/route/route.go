package route

import (
	"example.com/m/v2/websocket/v2/ws"
	"github.com/gin-gonic/gin"
)

func RouterApi(router *gin.Engine) {
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "go 你好")
	})
	router.LoadHTMLGlob("views/*")
	router.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"title": "HTML 模板渲染样例",
			"body":  "这里是内容",
		})
	})
	//	socket 服务
	router.GET("/socket", ws.Server)
}
