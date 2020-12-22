package conf

import "github.com/gorilla/websocket"



const archiveSize = 20
const chanSize = 10

const msgJoin = "[加入房间]"
const msgLeave = "[离开房间]"
const msgTyping = "[正在输入]"

const(
	EventTypeJoin = "event-join"
	EventTypeMsg = "event-message"
	EventTypeLeave = "event-leave"
	EventTypeBroadcast = "event-broadcast"
	EventTypeSystem = "event-system"
	EventTypeImage  = "event-image"
)

type Event struct{
	Type string `json:"type"`
	User string `json:"user"`
	UserCount int `json:"userCount"`
	Timestamp int64 `json:"timestamp"`
	Content string `json:"content"`
}

type Client struct{
	Username string
	RecvCHN <-chan Event
	SendCHN chan Event
	LeaveCHN chan string
	Ws *websocket.Conn
}

type Room struct{
	Users map[string]*websocket.Conn
	UserCount int
}






