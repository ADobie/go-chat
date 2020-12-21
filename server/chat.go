package server

import "github.com/gin-gonic/gin"

type Recvmsg struct{
	Recieve string `json:"recieve"`
}

type Sendmsg struct{
	Send string `json:"send"`
}

func Handler(c *gin.Context) {
	var recv Recvmsg
	var send Sendmsg
	conn,err := Websocket.upgrader.Upgrade(c.Writer,c.Request,nil)
    if err != nil{
		panic(err)
	}
	NewMessage:=make(chan string)
	go func() {
	for {
		err = conn.ReadJSON(&recv)
		if err != nil {
			close(NewMessage)
			return
		}
        NewMessage<-recv.Recieve
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
		err=conn.WriteJSON(&send)
		if err != nil {
			return
		}
	}
}

}