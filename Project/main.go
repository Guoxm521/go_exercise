package main

import (
	"example.com/m/v2/route"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	_config, _err := Init()
	if _err != nil {
		panic("配置错误" + _err.Error())
	}
	router := gin.Default()
	router.Use(Cors())
	route.RouterApi(router)
	runHost = strings.Join([]string{_config.Server.Host, ":", _config.Server.Port}, "")
	_gin := &http.Server{
		Addr:    runHost,
		Handler: router,
	}
	//_search := &spider.SearchStruct{
	//	Since:        "daily",
	//	SinceType:    1,
	//	Language:     "PHP",
	//	LanguageType: 1,
	//}
	//github := new(spider.GithubTrending)
	//github.NewCollector(_search).SpiderGithub()
	router.Run(_gin.Addr)
}
