package Controllers

import (
	"tarea1/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var cliente []Models.Cliente
	err := Models.GetAllUsers(&cliente)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cliente)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var cliente Models.Cliente
	c.BindJSON(&cliente)
	err := Models.CreateUser(&cliente)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cliente)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id_cliente := c.Params.ByName("id_cliente")
	var cliente Models.Cliente
	err := Models.GetUserByID(&cliente, id_cliente)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cliente)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var cliente Models.Cliente
	id_cliente := c.Params.ByName("id_cliente")
	err := Models.GetUserByID(&cliente, id_cliente)
	if err != nil {
		c.JSON(http.StatusNotFound, cliente)
	}
	c.BindJSON(&cliente)
	err = Models.UpdateUser(&cliente, id_cliente)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cliente)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var cliente Models.Cliente
	id_cliente := c.Params.ByName("id_cliente")
	err := Models.DeleteUser(&cliente, id_cliente)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id_cliente: "is deleted"})
	}
}