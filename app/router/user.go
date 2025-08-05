package router

import (
	"github.com/gin-gonic/gin"
	"go-book-center/app/middleware"
	"go-book-center/app/service"
)

func InitUserRouter(group *gin.RouterGroup) {
	group.Group("")

	group.GET("/user/:id", service.FindUserById)
	group.POST("/user/login", service.LoginUser)

	group.Use(middleware.AuthSession())
	{
		group.GET("/user/logout", service.LogoutUser)
	}
}
