package controllers

import (
	"net/http"

	"tarea1/models"

	"github.com/gin-gonic/gin"
)

type LoginClientInput struct {
	Id_cliente int    `json:"id_cliente" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}

func LoginClient(c *gin.Context) {
	var cliente models.Cliente
	var input LoginClientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	if err := models.DB.Where("id_cliente = ?", input.Id_cliente).Where("Contrasena = ?", input.Contrasena).First(&cliente).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no funca"})
		return
	} else {
		//err := models.DB.Where("id_cliente = ?", input.Id_cliente).First(&cliente)
		//fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"acceso valido": true})
	}
}
