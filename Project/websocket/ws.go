package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

//单个websocket消息
type Client struct {
	Id      string
	Group   string
	Socket  *websocket.Conn
	Message chan []byte
}

// messageData 单个发送数据信息
type MessageData struct {
	Id, Group string
	Message   []byte
}

// groupMessageData 组广播数据信息
type GroupMessageData struct {
	Group   string
	Message []byte
}

// 广播发送数据信息
type BroadCastMessageData struct {
	Message []byte
}

type Msg struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

const (
	msgTypeOnline        = 1 // 上线
	msgTypeOffline       = 2 // 离线
	msgTypeGetOnlineUser = 4 // 获取用户列表
	msgTypePrivateChat   = 5 // 私聊
	msgTypePublicChat    = 6 // 群聊
)

// 读信息，从 websocket 连接直接读取数据
func (c *Client) Read() {
	defer func() {
		WebsocketManager.UnRegister <- c
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}()
	for {
		messageType, message, err := c.Socket.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			break
		}
		log.Printf("client [%s] receive message: %s", c.Id, string(message))
		_message, status, _ := formatServeMsgStr(message)
		switch status {
		case msgTypePrivateChat:
			fmt.Println("私聊")
		case msgTypePublicChat:
			group, _ := clientMsg.Data.(map[string]interface{})["group"].(string)
			fmt.Println("group", group)
			WebsocketManager.SendGroup("123123", _message)
		default:
			if string(message) == "heartbeat" {
				c.Socket.WriteMessage(websocket.TextMessage, []byte(`{"status":0,"data":"heartbeat ok312321312312"}`))
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}()
	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("client [%s] write message: %s", c.Id, string(message))
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("client [%s] writemessage err: %s", c.Id, err)
			}
		}
	}
}

var clientMsg = Msg{}

func formatServeMsgStr(message []byte) ([]byte, int, error) {
	_err := json.Unmarshal(message, &clientMsg)
	if _err != nil {
		log.Printf("消息解析失败,消息[%s],错误[%s],", _err, string(message))
		return nil, 0, _err
	}
	username, _ := clientMsg.Data.(map[string]interface{})["username"].(string)
	uid, _ := clientMsg.Data.(map[string]interface{})["uid"].(string)
	//group, _ := clientMsg.Data.(map[string]interface{})["group"].(string)
	content, _ := clientMsg.Data.(map[string]interface{})["content"].(string)
	data := map[string]interface{}{
		"username": username,
		"uid":      uid,
		"group":    "123123",
		"content":  content,
		"time":     time.Now().UnixNano() / 1e6,
	}
	status := clientMsg.Status
	switch status {
	case msgTypePrivateChat:
		fmt.Println("私聊")
	case msgTypePublicChat:
		fmt.Println("群聊")
	default:
		fmt.Println("哈哈")
	}
	jsonStrServeMsg := Msg{
		Status: status,
		Data:   data,
	}
	serveMsgStr, _ := json.Marshal(jsonStrServeMsg)
	return serveMsgStr, status, nil
}
