package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

func NewEvent(typ string,username string,txt string) Event {
	return Event{
		Type: typ,
		User: username,
		Timestamp: time.Now().UnixNano()/1e6,
		Content: txt,
	}
}

func NewUser(username string) Client {
	tmp:=make(chan Client)
	s:=<-tmp
	s.Username = username
	return s
}

func(r *Room) AddToUsers(client Client){
	r.Users[client.Username] = client.Ws
}

func EstWebsocket(c *gin.Context) *websocket.Conn{
	conn,err := Websocket.Upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil{
		panic(err)
	}
	return conn
}

func NewRoom() *Room {
	return &Room{
		Users :map[string] *websocket.Conn{},
		UserCount: 0,
	}
}