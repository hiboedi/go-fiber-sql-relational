package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	dsn := "root@tcp(localhost:3306)/golang_test"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Succed connect to db")
}
