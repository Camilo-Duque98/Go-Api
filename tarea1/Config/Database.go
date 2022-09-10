package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)
var DB *gorm.DB
//Este archivo contiene la configuracion de la base de datos
type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
}

func BuildDBConfig() *DBConfig{
	dbConfig := DBConfig{
		Host: "localhost",
		Port: 3306,
		User: "root",
		Password: "123456789",
		DBName: "first_go",
	}
	fmt.Println(&dbConfig)
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string{
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}