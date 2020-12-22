package server

import (
	"github.com/gin-gonic/gin"
	"go-chat/conf"
)


type Recvmsg struct{
	Receive string `json:"receive"`
}

type Sendmsg struct{
	Send string `json:"send"`
}

func Handler(c *gin.Context) {
	var recv Recvmsg
	var send Sendmsg
	room:=conf.NewRoom()
	name:=getUsername(c)
	me:=conf.NewUser(name)
	me.Ws = conf.EstWebsocket(c)
	room.AddToUsers(me)
	//conn,err := Websocket.Upgrader.Upgrade(c.Writer,c.Request,nil)
    //if err != nil{
		//panic(err)
	//}
	NewMessage:=make(chan string)
	go func() {
	for {
		err := me.Ws.ReadJSON(&recv)
		if err != nil {
			close(NewMessage)
			return
		}
        NewMessage<-recv.Receive
	}
}()
for{
	select {
	case msg,ok:=<-NewMessage:
		if !ok{
			return
		}
		send=Sendmsg{
			Send:msg,
		}
		for _,v:=range room.Users{
			err:=v.WriteJSON(&send)
			if err != nil {
				return
			}
		}
	}
}

}