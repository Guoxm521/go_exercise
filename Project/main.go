package main

import "example.com/m/v2/Project/spider"

func main() {
	github := new(spider.GithubTrending)
	github.NewCollector().SpiderGithub()
	//_config,_err := Init()
	//if _err != nil {
	//	panic("配置错误" + _err.Error())
	//}
	//router := gin.Default()
	//route.RouterApi(router)
	//runHost = strings.Join([]string{_config.Server.Host, ":", _config.Server.Port}, "")
	//_gin := &http.Server{
	//	Addr:    runHost,
	//	Handler: router,
	//}
	//router.Run(_gin.Addr)
}
