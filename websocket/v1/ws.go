package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func NewRouter() *gin.Engine {
	server := gin.Default()
	server.Use(Cors())
	socket := RunSocket
	group := server.Group("")
	{
		group.GET("/socket", socket)
	}
	return server
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		c.Next()
	}
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(c *gin.Context) {
	// 可以接受socket参数
	user := c.Query("user")
	if user == "" {
		return
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	_, message, err := ws.ReadMessage()
	if nil != err {
		fmt.Println(err.Error())
	}
	ws.WriteMessage(websocket.TextMessage, message)
	for {
		time.Sleep(3 * time.Second)
		ws.WriteMessage(1, []byte("31231232131231232312"))
	}
}
