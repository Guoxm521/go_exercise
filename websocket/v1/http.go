package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpServer() {
	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "HTML 模板渲染样例",
			"body":  "这里是内容",
		})
	})
	engine.Run(":9999")
}
