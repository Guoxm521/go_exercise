package main

import (
	"example.com/m/v2/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	runHost   string
	AppConfig = new(database.ConfigModel)
)

func Init() (*database.ConfigModel, error) {
	_config, _err := database.ReadAppConfig()
	if _err != nil {
		return AppConfig, _err
	}
	AppConfig = _config
	database.ConnMySQL(AppConfig.MySql)
	gin.DisableConsoleColor()
	gin.SetMode(AppConfig.Server.Mode)
	return AppConfig, nil
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
