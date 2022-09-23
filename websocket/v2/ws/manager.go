package ws

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

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
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}
	go client.Read()
	go client.Write()
	time.Sleep(3 * time.Second)
	//测试函数
	//Demo(client.Message)
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

type Manager struct {
	Group        map[string]map[string]*Client
	Register     chan *Client
	UnRegister   chan *Client
	Message      chan *MessageData
	GroupMessage chan *GroupMessageData
}

var WebsocketManager = Manager{
	Group:        make(map[string]map[string]*Client),
	Register:     make(chan *Client, 128),
	UnRegister:   make(chan *Client, 128),
	GroupMessage: make(chan *GroupMessageData, 128),
	Message:      make(chan *MessageData, 128),
}

//处理单个  Client  发送数据
//处理单个  group  广播数据
// 向指定的 client 发送数据
// 向指定的 Group 广播
