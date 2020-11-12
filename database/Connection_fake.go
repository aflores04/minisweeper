package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewConnectionTest() DbConnection {
	return TestConnection{}
}

type TestConnection struct {}

func (c TestConnection) Connect() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"root",
		"localhost",
		"minisweeper",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("error in connection with db: ", err)
	}

	return db
}
