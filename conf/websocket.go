package conf

import (
	"github.com/gorilla/websocket"
	"net/http"
)


var Websocket = &ws{
	Upgrader: &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	},
}

type ws struct {
	Upgrader *websocket.Upgrader
}
