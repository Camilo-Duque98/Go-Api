package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:123456789@tcp(localhost:3306)/tarea_1_sd")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Producto{})

	DB = database
}
