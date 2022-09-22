package main

import (
	"example.com/m/v2/websocket/v2/database"
	"example.com/m/v2/websocket/v2/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	_config, _err := database.ReadAppConfig()
	if _err != nil {
		fmt.Printf(" config file read failed: %s", _err)
		return
	}
	router := gin.Default()
	route.RouterApi(router)
	runHost := strings.Join([]string{_config.Server.Host, ":", _config.Server.Port}, "")
	_gin := &http.Server{
		Addr:    runHost,
		Handler: router,
	}
	router.Run(_gin.Addr)
}
