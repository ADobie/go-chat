package main

import (
	"github.com/gin-gonic/gin"
	"go-chat/server"
)


func main(){
	server.Database()
	r:=gin.Default()
	r.GET("/ws",func(c *gin.Context){
		server.Handler(c)
	})
	r.Run()
}

