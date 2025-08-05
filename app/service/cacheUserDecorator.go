package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-book-center/app/common"
	"go-book-center/app/database"
	"log"
	"net/http"
)

var logger = common.Logger

func CacheUserDecorator(h gin.HandlerFunc, porm string, readKeyPattern string, user interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyId := c.Param(porm)
		redisKey := fmt.Sprintf("%s_%s", readKeyPattern, keyId)
		data, err := database.Redis.Get(c, redisKey).Result()
		if err != nil {
			// call service.RedisOneUser, get user from db
			h(c)
			dbRes, exists := c.Get(redisKey)
			if !exists {
				dbRes = user
			}
			//set user in redis
			redisData, _ := json.Marshal(dbRes)
			err := database.Redis.Set(c, redisKey, string(redisData), 0).Err()
			if err != nil {
				logger.Errorf("cacheUserDecorator: redis set error: %v", err)
			}
			c.JSON(http.StatusOK, gin.H{"msg": "success from db", "data": dbRes})
			return
		}
		log.Printf("cacheUserDecorator: redis get data: %v", data)
		json.Unmarshal([]byte(data), &user)
		c.JSON(http.StatusOK, gin.H{"msg": "success from redis", "data": user})
	}
}
