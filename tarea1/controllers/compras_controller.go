package controllers

import (
	//"fmt"

	"fmt"
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

type UpdateProducts struct {
	Nombre          string `json:"nombre" binding:"required"`
	Cantidad        int    `json:"cantidad_disponible" binding:"required"`
	Precio_unitario int    `json:"precio_unitario" binding:"required"`
}

func CantidadData(ID_PRODUCTO int) int {
	var products models.Producto

	if err := models.DB.Where("id_producto = ?", ID_PRODUCTO).First(&products).Error; err != nil {
		fmt.Println(err)
		return -1
	}

	return products.Cantidad_disponible
}
func UpdateProductIncrease(ID_PRODUCTO int, CantidadProductos int, CantidadCompra int) int {
	var products models.Producto
	var updateProduct UpdateProducts

	if err := models.DB.Where("id_producto = ?", ID_PRODUCTO).First(&products).Error; err != nil {
		fmt.Println(err)
		return 0
	}
	//updateProduct.Nombre = products.Nombre
	//updateProduct.Precio_unitario = products.Precio_unitario
	//updateProduct.Cantidad = CantidadProductos - CantidadCompra
	lista := UpdateProductInput{products.Nombre, CantidadProductos - CantidadCompra, products.Precio_unitario}
	fmt.Println(updateProduct)
	//models.DB.Model(&products).Where("id_producto = ?", ID_PRODUCTO).Updates(updateProduct)
	models.DB.Model(&products).Where("id_producto = ?", ID_PRODUCTO).Updates(lista)
	return products.Cantidad_disponible
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
				//c.JSON(http.StatusBadRequest, gin.H{"error": "Producto no encontrado"})
				c.JSON(http.StatusBadRequest, gin.H{"error": array.Id_producto})
				return
			} else {
				fmt.Println(array.Id_producto)
				returnFunctionUpdate := CantidadData(array.Id_producto)
				fmt.Println(returnFunctionUpdate)
				if array.Cantidad <= returnFunctionUpdate {
					returnUpdateProductIncrease := UpdateProductIncrease(array.Id_producto, returnFunctionUpdate, array.Cantidad)
					fmt.Println(returnUpdateProductIncrease)
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "productos insuficientes"})
				}
				detalle := models.Detalle{Id_compra: CompraID, Id_producto: array.Id_producto, Cantidad: array.Cantidad}
				models.DB.Create(&detalle)
			}

		}

		c.JSON(http.StatusOK, gin.H{"id_compra": compra.Id_compra})
	}

}
