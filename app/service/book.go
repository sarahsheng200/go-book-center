package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-book-center/app/middleware"
	"go-book-center/app/repository"
	"go-book-center/app/schema"
	"net/http"
	"strconv"
)

func AddBook(c *gin.Context) {

	book := schema.Book{}
	userid := middleware.GetSession(c)
	fmt.Println("=======", userid)
	middleware.AuthSession()
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook := repository.AddBook(book)
	c.JSON(http.StatusOK, gin.H{"data": newBook})
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("id")
	book := repository.GetBookInfo(bookId)
	categoryId := strconv.FormatInt(book.CategoryId, 10)
	book.CategoryChain = repository.GetAllCategories(categoryId)
	book.Category = book.CategoryChain[0]

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetAllBooks(c *gin.Context) {
	books := repository.GetBookList()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("id")
	isEffected := repository.DeleteBook(bookId)
	if !isEffected {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book Not Found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": isEffected})
}

func UpdateBook(c *gin.Context) {
	bookId := c.Param("id")
	book := schema.UpdateBook{}
	if bookId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book id can't be empty"})
	}
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json format is wrong: " + err.Error()})
		return
	}
	if isExisted := repository.IsExistedBook(bookId); !isExisted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book Not Found"})
	}

	latestBook := repository.UpdateBook(bookId, book)
	c.JSON(http.StatusOK, gin.H{"data": latestBook})
}
