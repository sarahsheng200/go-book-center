package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"go-book-center/app/common"
	"go-book-center/app/config"
	"go-book-center/app/schema"
	"net/http"
)

type Context struct {
	Ctx *gin.Context
}

var conf = config.Conf.Session

var loggers = common.Logger

func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(conf.StoreKey))
	return sessions.Sessions(conf.Name, store)
}

func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user schema.UserAuth
		session := sessions.Default(c)
		data, ok := session.Get(conf.SessionKey).(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "please login"})
			c.Abort()
			return
		}
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "please login, json format is wrong"})
		}
		c.Set(conf.Name, user)
		c.Next()
	}
}

func CheckIsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, exists := c.Get(config.Conf.Session.Name)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"msg": "checkIsAdmin: no context session found, please login"})
			c.Abort()
			return
		}
		user, ok := data.(schema.UserAuth)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "checkIsAdmin:  context session format error"})
			c.Abort()
			return
		}
		if user.IsAdmin != 1 {
			c.JSON(http.StatusForbidden, gin.H{"msg": "checkIsAdmin: this user is not admin"})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "checkIsAdmin: ok"})
	}
}

func SaveSession(c *gin.Context, user string) {
	session := sessions.Default(c)
	session.Set(conf.SessionKey, user)
	err := session.Save()
	if err != nil {
		loggers.Errorf("Logger: save session error:", err)
	}
}

func GetSession(c *gin.Context) interface{} {
	session := sessions.Default(c)
	data := session.Get(conf.SessionKey)

	return data
}

func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		loggers.Errorf("Logger: clear session error:", err)
	}
}
