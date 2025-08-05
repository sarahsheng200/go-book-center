package database

import (
	con "go-book-center/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DBconn *gorm.DB
var err error

func MysqlConnection() {
	config := con.Conf.Mysql
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Path + ")/" + config.Database + "?" + config.Config
	DBconn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: config.SingularTable, // 启用单数表名
		},
	})
	if err != nil {
		log.Fatal(err)
	}

}
