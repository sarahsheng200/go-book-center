package router

import (
	"github.com/gin-gonic/gin"
	conf "go-book-center/app/config"
	"go-book-center/app/middleware"
)

func InitRouter() *gin.Engine {
	config := conf.Conf.Server

	gin.SetMode(config.Mode)
	router := gin.New()
	router.Use(
		middleware.SetSession(),
		gin.Recovery(),
		middleware.LoggerToFile(),
	)

	group := router.Group(config.UrlPrefix)

	InitBookRouter(group)
	InitUserRouter(group)
	InitCategoryRouter(group)

	return router

}
