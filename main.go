package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-chat/server"
)

func main(){
	server.Database()
	r:=gin.Default()
	store := cookie.NewStore([]byte("username"))
	r.Use(sessions.Sessions("session", store))
	r.POST("/login", func(c *gin.Context) {
		server.Login(c)
	})
	r.POST("/register", func(c *gin.Context) {
		server.Register(c)
	})
	r.GET("/ws",func(c *gin.Context){
		if server.GetSession(c){
			server.Handler(c)
		}
	})
	r.Run()
}

