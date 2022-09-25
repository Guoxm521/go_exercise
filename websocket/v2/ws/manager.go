package ws

import (
	"log"
	"sync"
)

type Manager struct {
	Group                   map[string]map[string]*Client
	Register                chan *Client
	UnRegister              chan *Client
	Message                 chan *MessageData
	GroupMessage            chan *GroupMessageData
	Lock                    sync.Mutex
	groupCount, clientCount uint
}

var WebsocketManager = Manager{
	Group:        make(map[string]map[string]*Client),
	Register:     make(chan *Client, 128),
	UnRegister:   make(chan *Client, 128),
	GroupMessage: make(chan *GroupMessageData, 128),
	Message:      make(chan *MessageData, 128),
	groupCount:   0,
	clientCount:  0,
}

// 启动 websocket 管理器
func (manager *Manager) Start() {
	log.Printf("websocket manage start")
	for {
		select {
		//注册
		case client := <-manager.Register:
			log.Printf("client [%s] connect", client.Id)
			log.Printf("register client [%s] to group [%s]", client.Id, client.Group)
			manager.Lock.Lock()
			if manager.Group[client.Group] == nil {
				manager.Group[client.Group] = make(map[string]*Client)
				manager.groupCount += 1
			}
			manager.Group[client.Group][client.Id] = client
			manager.clientCount += 1
			manager.Lock.Unlock()
		//注销
		case client := <-manager.UnRegister:
			log.Printf("unregister client [%s] from group [%s]", client.Id, client.Group)
			manager.Lock.Lock()
			if _, ok := manager.Group[client.Group]; ok {
				if _, ok := manager.Group[client.Group][client.Id]; ok {
					close(client.Message)
					delete(manager.Group[client.Group], client.Id)
					manager.clientCount -= 1
					if len(manager.Group[client.Group]) == 0 {
						delete(manager.Group, client.Group)
						manager.groupCount -= 1
					}
				}
			}
			manager.Lock.Unlock()
		}
	}
}

//处理单个  Client  发送数据
func (manager *Manager) SendService() {
	for {
		select {
		case data := <-manager.Message:
			if groupMap, ok := manager.Group[data.Group]; ok {
				if conn, ok := groupMap[data.Id]; ok {
					conn.Message <- data.Message
				}
			}
		}
	}
}

//处理单个  group  广播数据
// 向指定的 client 发送数据
// 向指定的 Group 广播
//注册
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

//注销
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

// 当前组个数
func (manager *Manager) LenGroup() uint {
	return manager.groupCount
}

// 当前连接个数
func (manager *Manager) LenClient() uint {
	return manager.clientCount
}

// 获取 wsManager 管理器信息
func (manager *Manager) Info() map[string]interface{} {
	managerInfo := make(map[string]interface{})
	managerInfo["groupLen"] = manager.LenGroup()
	managerInfo["clientLen"] = manager.LenClient()
	managerInfo["chanRegisterLen"] = len(manager.Register)
	managerInfo["chanUnregisterLen"] = len(manager.UnRegister)
	managerInfo["chanMessageLen"] = len(manager.Message)
	managerInfo["chanGroupMessageLen"] = len(manager.GroupMessage)
	return managerInfo
}
