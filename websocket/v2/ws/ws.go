package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
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

func (c *Client) Read() {
	for {
		messageType, message, err := c.Socket.ReadMessage()
		fmt.Println("message", message)
		if err != nil || messageType == websocket.CloseMessage {
			fmt.Println("err", err)
			fmt.Println("messageType", websocket.CloseMessage)
			c.Socket.Close()
			close(c.Message)
			break
		}
		log.Printf("client [%s] receive message: %s", "id", string(message))
		c.Message <- message
	}
}

func (c *Client) Write() {
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
