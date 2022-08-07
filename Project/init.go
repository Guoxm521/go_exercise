package main

import (
	"example.com/m/v2/database"
	"github.com/gin-gonic/gin"
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
