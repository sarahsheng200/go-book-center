package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-book-center/app/config"
	"go-book-center/app/database"
	"go-book-center/app/middleware"
	"go-book-center/app/repository"
	"go-book-center/app/schema"
	"net/http"
)

func FindUserById(c *gin.Context) {
	user := schema.User{}
	id := c.Param("id")
	if id == "0" {
		c.JSON(http.StatusNotFound, "Error")
	}
	isUseRedis := config.Conf.Server.UseRedis

	if isUseRedis {
		redisKey := fmt.Sprintf("user_%s", id)
		data, err := database.Redis.Get(c, redisKey).Result()
		if err != nil {
			// there is no redis data
			user = repository.FindUserById(id)
			redisUser, _ := json.Marshal(user)
			database.Redis.Set(c, redisKey, string(redisUser), 0)
			c.JSON(http.StatusOK, gin.H{"msg": "success from db", "data": user})
			return
		}

		if err := json.Unmarshal([]byte(data), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "json format is wrong", "error": err})
		}

		c.JSON(http.StatusOK, gin.H{"msg": "success from redis", "data": user})
	} else {
		user = repository.FindUserById(id)
		c.JSON(http.StatusOK, gin.H{"msg": "success from db and userRedis is off", "data": user})
	}
}

func LoginUser(c *gin.Context) {
	accountNum := c.PostForm("accountNum")
	password := c.PostForm("password")
	user := repository.CheckUserPassword(accountNum, password)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	userSession := gin.H{
		"id":      user.Id,
		"name":    user.Name,
		"isAdmin": user.IsAdmin,
	}
	value, err := json.Marshal(userSession)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}
	middleware.SaveSession(c, string(value))
	c.JSON(http.StatusOK, gin.H{"msg": "Login success!", "user": user, "session": middleware.GetSession(c)})
}

func LogoutUser(c *gin.Context) {
	middleware.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{"msg": "Logout success!"})
}
