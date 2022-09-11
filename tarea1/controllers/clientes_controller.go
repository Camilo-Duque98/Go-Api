package controllers

import (
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
 	"tarea1/models"
)

type CreateClientInput struct {
	Nombre string `json:"nombre" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}

type LoginClientInput struct {
	Id_cliente int `json:"id_cliente" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}

func LoginClient(c *gin.Context){
	var cliente models.Cliente
	//nombre := c.Request.URL.Query().Get("nombre")
	var input LoginClientInput
	fmt.Println(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	if err := models.DB.Where("id_cliente = ?", input.Id_cliente).First(&cliente).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"no funca"})
		return 
	} else{

		c.JSON(http.StatusOK, gin.H{"acceso valido": true})
	}
}

/*func FindClient(c *gin.Context){
	var cliente models.Cliente
	id_cliente := c.Request.URL.Query().Get("id_cliente")
	if err := models.DB.Where("id_cliente", id_cliente).Find(&cliente).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H("data":cliente))
}*/