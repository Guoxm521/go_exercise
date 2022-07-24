package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"Message":"docker部署镜像你好呀",
		})
	})
	r.GET("/go/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"Message":"docker部署镜像你好呀 go",
		})
	})
	r.Run()
}
