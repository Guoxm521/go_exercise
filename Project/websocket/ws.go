package websocket

import (
	"encoding/json"
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

var clientMsg = Msg{}

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
		//c.Message <- message
		WebsocketManager.SendGroup("123123", message)
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
			err := c.Socket.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				log.Printf("client [%s] writemessage err: %s", c.Id, err)
			}
		}
	}
}

//avatar_id
//room_id
//to_user
//uid
//username
func formatServeMsgStr(status int) ([]byte, Msg) {
	data := map[string]interface{}{
		"username": clientMsg.Data.(map[string]interface{})["username"].(string),
		"uid":      clientMsg.Data.(map[string]interface{})["uid"].(string),
		"group":    "123123",
		"time":     time.Now().UnixNano() / 1e6,
	}
	if status == msgTypePrivateChat || status == msgTypePublicChat {
		data["content"] = clientMsg.Data.(map[string]interface{})["content"].(string)
	}
	jsonStrServeMsg := Msg{
		Status: status,
		Data:   data,
	}
	serveMsgStr, _ := json.Marshal(jsonStrServeMsg)
	return serveMsgStr, jsonStrServeMsg
}
