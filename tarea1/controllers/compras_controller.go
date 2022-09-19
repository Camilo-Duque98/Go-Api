package controllers

import (
	"fmt"

	"net/http"
	"tarea1/models"

	"github.com/gin-gonic/gin"
)

type CreateCompraInput struct {
	Id_cliente int `json:"id_cliente" binding:"required"`
}

/*type Detalle struct {
	Id_compra   int `json:"id_compra" binding:"required"`
	Id_producto int `json:"id_producto" binding:"required"`
	Cantidad    int `json:"cantidad" binding:"required"`
}*/

func CreateCompra(c *gin.Context) {
	var cliente models.Cliente
	var input CreateCompraInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	if err := models.DB.Where("id_cliente = ?", input.Id_cliente).First(&cliente).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no funca"})
		return
	} else {
		compra := models.Compra{Id_cliente: input.Id_cliente}
		models.DB.Create(&compra)
		//models.DB.Model(&compra).Association("cliente")
		c.JSON(http.StatusOK, gin.H{"id_compra": compra.Id_compra})
		fmt.Println(compra.Id_compra)
	}
}

/*func FindCompras(c *gin.Context) {
	var compra []models.Compra
	models.DB.Find(&compra)
	models.DB.Model(&compra).Association("id_cliente")
	c.JSON(http.StatusOK, gin.H{"data": compra})

}*/
