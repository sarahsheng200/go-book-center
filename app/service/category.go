package service

import (
	"github.com/gin-gonic/gin"
	"go-book-center/app/repository"
	"net/http"
)

func GetCategoryChain(c *gin.Context) {
	categoryId := c.Param("id")

	categoryChain := repository.GetAllCategories(categoryId)

	if len(categoryChain) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categoryChain})
}

func GetCategoryById(c *gin.Context) {
	categoryId := c.Param("id")
	category := repository.GetCategoryById(categoryId)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func GetCategoryNameList(c *gin.Context) {
	var categoryNameList []string
	categoryId := c.Param("id")
	categoryChain := repository.GetAllCategories(categoryId)
	if len(categoryChain) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category Not Found"})
		return
	}
	for _, category := range categoryChain {
		categoryNameList = append(categoryNameList, category.Name)
	}
	c.JSON(http.StatusOK, gin.H{"data": categoryNameList})
}
