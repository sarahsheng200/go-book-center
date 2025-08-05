package repository

import (
	"go-book-center/app/database"
	"go-book-center/app/schema"
	"log"
	"strconv"
)

func GetCategoryById(id string) schema.Category {
	var category schema.Category

	database.DBconn.First(&category, id)
	return category
}

func GetAllCategories(id string) schema.Categories {
	if id == "0" {
		log.Println("category_repo: Category Id is invalid")
		return nil
	}
	var category schema.Category
	var categoryChain []schema.Category

	for {
		category = GetCategoryById(id)
		categoryChain = append(categoryChain, category)
		if category.ParentId == nil {
			break
		}
		id = strconv.FormatInt(*category.ParentId, 10)
		logger.Infof("category: %v", categoryChain)
	}

	return categoryChain
}
