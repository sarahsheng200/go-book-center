package repository

import (
	"go-book-center/app/common"
	"go-book-center/app/database"
	"go-book-center/app/schema"
)

var logger = common.Logger

func GetBookInfo(id string) schema.Book {
	var book schema.Book
	database.DBconn.Joins("LEFT JOIN author ON author.id = book.author_id").
		Joins("LEFT JOIN country ON country.id = author.country_id").
		Where("id = ? AND is_deleted = 0", id).
		Preload("Author.Country").
		First(&book, id)
	return book
}

func IsExistedBook(id string) bool {
	var book schema.Book
	database.DBconn.Where("id = ? AND is_deleted = 0", id).First(&book)
	return book.Id > 0
}

func GetBookList() schema.BookList {
	var booklist schema.BookList
	database.DBconn.Find(&booklist)
	return booklist
}

func AddBook(book schema.Book) schema.Book {
	database.DBconn.Create(&book)
	return book
}

func DeleteBook(id string) bool {
	book := schema.Book{}
	book.IsDeleted = 1
	rowsAffected := database.DBconn.Where("id = ? AND is_deleted = 0", id).Updates(&book).RowsAffected
	return rowsAffected > 0
}

func UpdateBook(id string, book schema.UpdateBook) schema.UpdateBook {
	database.DBconn.Model(&book).Where("id=?", id).Updates(&book)
	return book
}
