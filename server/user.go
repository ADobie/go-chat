package server

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Confirm  string `json:"confirm" gorm:"-"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Username string `json:"username" gorm:"type:varchar(255)"`
}

type Msg struct {
	gorm.Model
	Message  string `json:"message" gorm:"type:varchar(255)"`
	Username string `json:"username" gorm:"type:varchar(255)"`
}

func register(c *gin.Context) {
	var regData User
	err := c.BindJSON(&regData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "Msg": "注册失败！"})
		return
	}
	if regData.Password == regData.Confirm && regData.Password != "" {
		q := Db.Where("username = ?", regData.Username).Find(&regData).RowsAffected
		if q != 0 {
			c.JSON(200, gin.H{"code": "400", "Msg": "用户名已存在！"})
			return
		}
		Db.Create(&regData)
		c.JSON(200, gin.H{"code": "200", "Msg": "注册成功！"})
	}
}

func login(c *gin.Context) {
	var logData User
	err := c.BindJSON(&logData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "Msg": "请求出错！"})
		return
	}
	q := Db.Where("username=? AND password=?", logData.Username, logData.Password).Find(&logData).RowsAffected
	if q > 0 {
		session := sessions.Default(c)
		session.Set("username", logData.Username)
		session.Save()
		c.JSON(200, gin.H{"code": "200", "Msg": "登录成功！"})
	} else {
		c.JSON(200, gin.H{"code": "400", "Msg": "用户名或密码错误！"})
		return
	}
}

func getSession(c *gin.Context) bool {
	session := sessions.Default(c)
	username := session.Get("username")
	if username != nil {
		return true
	} else {
		return false
	}
}

func getUsername(c *gin.Context) string {
	var tmpData User
	session := sessions.Default(c)
	name := session.Get("username")
	err := Db.Where("username=?", name).Find(&tmpData).Error
	if err != nil {
		fmt.Println(err)
	}
	return tmpData.Username
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()
	session.Clear()
	c.Redirect(302, "/")
}
