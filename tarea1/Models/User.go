package Models

import (
	"tarea1/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(cliente *[]Cliente) (err error) {
	if err = Config.DB.Find(cliente).Error; err != nil {
		return err
	}
	fmt.Println(cliente)
	return nil
}

//CreateUser ... Insert New data
func CreateUser(cliente *Cliente) (err error) {
	if err = Config.DB.Create(cliente).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(cliente *Cliente, id_cliente string) (err error) {
	if err = Config.DB.Where("id_cliente = ?", id_cliente).First(cliente).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(cliente *Cliente, id_cliente string) (err error) {
	fmt.Println(cliente)
	Config.DB.Save(cliente)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(cliente *Cliente, id_cliente string) (err error) {
	Config.DB.Where("id_cliente = ?", id_cliente).Delete(id_cliente)
	return nil
}