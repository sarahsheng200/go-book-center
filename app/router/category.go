package router

import (
	"github.com/gin-gonic/gin"
	"go-book-center/app/service"
)

func InitCategoryRouter(group *gin.RouterGroup) {
	group.Group("")

	group.GET("/category/chain/:id", service.GetCategoryChain)
	group.GET("/category/:id", service.GetCategoryById)
	group.GET("/category/names/:id", service.GetCategoryNameList)
}
