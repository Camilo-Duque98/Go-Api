package controllers

import (
	//"fmt"

	"net/http"
	"tarea1/models"

	"github.com/gin-gonic/gin"
)

type CreateCompraInput struct {
	Id_cliente    int            `json:"id_cliente" binding:"required"`
	DetalleInputs []DetalleInput `json:"productos" binding:"required"`
}

type DetalleInput struct {
	Id_producto int `json:"id_producto" binding:"required"`
	Cantidad    int `json:"cantidad" binding:"required"`
}

func CreateCompra(c *gin.Context) {

	var cliente models.Cliente
	var producto models.Producto
	var input CreateCompraInput

	var CompraID int
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
		CompraID = compra.Id_compra

		for _, array := range input.DetalleInputs {
			if err := models.DB.Where("id_producto = ?", array.Id_producto).First(&producto).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Producto no encontrado"})
				return
			} else {
				detalle := models.Detalle{Id_compra: CompraID, Id_producto: array.Id_producto, Cantidad: array.Cantidad}
				models.DB.Create(&detalle)
			}

		}

		c.JSON(http.StatusOK, gin.H{"id_compra": compra.Id_compra})
	}

}
