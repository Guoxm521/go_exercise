package ws

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func Server(gin *gin.Context) {
	go WebsocketManager.Start()
	go WebsocketManager.SendService()
	RunSocket(gin)
}

func RunSocket(gin *gin.Context) {
	wsUpgrade := websocket.Upgrader{}
	wsUpgrade.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := wsUpgrade.Upgrade(gin.Writer, gin.Request, nil)
	defer conn.Close()
	if err != nil {
		log.Printf("websocket connect error: %s", gin.Param("channel"))
		return
	}
	client := &Client{
		Id:      uuid.NewV4().String(),
		Group:   gin.Param("channel"),
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}
	WebsocketManager.RegisterClient(client)
	go client.Read()
	go client.Write()
	//测试函数
	//Demo(client.Message)
	for {
		time.Sleep(1 * time.Second)
		client.Message <- []byte("测试")
	}

}

func Demo(message chan []byte) {
	path, _ := os.Getwd()
	fmt.Println(path + "\\shell_demo.sh")
	cmd := exec.Command("D:/Program Files (x86)/Git/bin/bash", path+"\\shell_demo.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			fmt.Println(err2)
			break
		}
		message <- []byte(line)
		fmt.Println(line)
	}

	cmd.Wait()
}
cd