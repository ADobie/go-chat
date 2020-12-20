package main

import (
	"github.com/gin-gonic/gin"
	"go-chat/server"
)


func main(){
	server.Database()
	r:=gin.Default()
	r.Run()
}

