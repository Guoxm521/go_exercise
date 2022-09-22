package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func runSocket(gin *gin.Context) {
	wsUpgrade := websocket.Upgrader{}
	wsUpgrade.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	c, _ := wsUpgrade.Upgrade(gin.Writer, gin.Request, nil)
	defer c.Close()
	go read(c)
	go write(c)
}

func read(c *websocket.Conn) {

}

func write(c *websocket.Conn) {

}
