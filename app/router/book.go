package router

import (
	"github.com/gin-gonic/gin"
	"go-book-center/app/middleware"
	"go-book-center/app/service"
)

func InitBookRouter(group *gin.RouterGroup) {
	book := group.Group("/book").Use(middleware.AuthSession())
	{
		book.GET("/", service.GetAllBooks)
		book.GET("/:id", service.GetBookById)

		book.Use(middleware.CheckIsAdmin())
		{
			book.POST("/more", service.AddBook)
			book.PUT("/:id", service.UpdateBook)
			book.DELETE("/:id", service.DeleteBook)
		}
	}

}
