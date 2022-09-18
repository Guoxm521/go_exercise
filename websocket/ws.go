package main

import "github.com/gorilla/websocket"

func main() {

}

const (
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs() {

}
