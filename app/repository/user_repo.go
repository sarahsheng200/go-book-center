package repository

import (
	"go-book-center/app/database"
	"go-book-center/app/schema"
)

func CheckUserPassword(num string, password string) schema.User {
	user := schema.User{}
	database.DBconn.Where("account_number=? and password=?", num, password).First(&user)
	return user
}

func FindUserById(id string) schema.User {
	var user schema.User
	database.DBconn.First(&user, id)
	//database.DBconnect.Where("id=?", id).First(&user)
	return user
}
