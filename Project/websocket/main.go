package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Server(gin *gin.Context) {
	go WebsocketManager.Start()
	go WebsocketManager.SendService()
	go WebsocketManager.SendGroupService()
	RunSocket(gin)
}

func RunSocket(gin *gin.Context) {
	defer func() {
		_err := recover()
		fmt.Println("_err12", _err)
	}()
	wsUpgrade := websocket.Upgrader{
		Subprotocols: []string{"token"},
	}
	wsUpgrade.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := wsUpgrade.Upgrade(gin.Writer, gin.Request, nil)
	if err != nil {
		log.Printf("websocket connect error: %s", gin.Param("channel"))
		return
	}
	client := &Client{
		Id: uuid.NewV4().String(),
		//Group:   gin.Param("channel"),
		Group:   "123123",
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}
	WebsocketManager.RegisterClient(client)
	go client.Read()
	go client.Write()
	client.Message <- []byte("测试")
	//测试函数
	//Demo(client.Message)
	//for {
	//	time.Sleep(120 * time.Second)
	//	client.Message <- []byte("测试3123")
	//}
}
