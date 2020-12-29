package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-chat/conf"
	"go-chat/server"
	"log"
	"net/http"
)

func main(){
	server.Database()
	room := conf.NewRoom()
	//_, err := server.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJhZG1pbiIsImV4cCI6MTYwOTAwMTg3MSwiaXNzIjoiZ2luLWJsb2cifQ.sAws8iZ0VMF4Hn0exp4URjLu-0B4SI2DTk3ZhZcPFT4")
	//if err != nil {
	//	panic(err)
	//}
	r:=gin.Default()
	r.Use(Cors())

	store := cookie.NewStore([]byte("username"))
	r.Use(sessions.Sessions("session", store))
	api:=r.Group("/api")
	api.POST("/register", func(c *gin.Context) {
		server.Register(c)
	})
	api.POST("/login", func(c *gin.Context) {
		server.Login(c)
	})
	a:=r.Group("/")
	a.Use(server.JWT())
	a.GET("/ws",func(c *gin.Context){
		//if server.GetSession(c){
			server.Handler(c,room)
		//}
	})
	r.Run(":4444")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}



