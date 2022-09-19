package controllers

import (
	//"fmt"
	"net/http"

	"tarea1/models"

	"github.com/gin-gonic/gin"
)

type CreateCompraInput struct {
	Id_cliente int `json:"id_cliente" binding:"required"`
}

func CreateCompra(c *gin.Context) {
	var input CreateCompraInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	compra := models.Compra{Id_cliente: input.Id_cliente}
	models.DB.Create(&compra)
	c.JSON(http.StatusOK, gin.H{"data": compra.Id_compra})
}
