package server

import (
	"github.com/gorilla/websocket"
	"net/http"
)


var Websocket = &ws{
	upgrader: &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	},
}

type ws struct {
	upgrader *websocket.Upgrader
}
