package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:emmanuelcudjoe90@tcp(127.0.0.1:3306)/blogApiV1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Unable to connect to database " + err.Error())
		return nil
	}

	log.Println(db.Name() + " connected sucessfully")

	return db
}
