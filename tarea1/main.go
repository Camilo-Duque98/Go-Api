package main

import (
	"tarea1/Config"
	"tarea1/Models"
	"tarea1/Routes"
	"fmt"
	

	"github.com/jinzhu/gorm"
   )
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", 
	Config.DbURL(Config.BuildDBConfig()))
	
	
	if err != nil {
		fmt.Println("Status:", err)
	}
	
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Cliente{})
	
		r := Routes.SetupRouter()
	//running
	fmt.Println("jsdakljdal")
	r.Run()
   }