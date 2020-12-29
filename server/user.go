package server

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	_ "github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
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

var jwtSecret = []byte("31415926535")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}


func GenerateToken(username string,password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				fmt.Println(err)
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}


func Register(c *gin.Context) {
	var regData User
	err := c.BindJSON(&regData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "Msg": "注册失败！"+err.Error()})
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

func Login(c *gin.Context) {
	var logData User
	err := c.BindJSON(&logData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "Msg": "请求出错！"})
		return
	}
	q := Db.Where("username=? AND password=?", logData.Username, logData.Password).Find(&logData).RowsAffected
	if q > 0 {
	token,err:=GenerateToken(logData.Username, logData.Password)
	if err!=nil{
		c.JSON(200,gin.H{"code":"400","Msg":"请求出错！"})
		return
	}else{
		c.JSON(200, gin.H{"code": "200", "Msg": "登录成功！","token":token})
	}
	}else {
		c.JSON(200, gin.H{"code": "400", "Msg": "用户名或密码错误！"})
		return
	}
}

func getUsername(c *gin.Context) string {
	token:=c.Query("token")
	claim, err :=ParseToken(token)
	if err !=nil {
		panic(err)
	}
	return claim.Username
}


