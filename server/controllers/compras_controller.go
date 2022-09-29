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

	if err := models.DB.Where("id_producto = ?", ID_PRODUCTO).First(&products).Error; err != nil {
		fmt.Println(err)
		return 0
	}
	lista := UpdateProductInput{products.Nombre, CantidadProductos - CantidadCompra, products.Precio_unitario}
	models.DB.Model(&products).Where("id_producto = ?", ID_PRODUCTO).Updates(lista)
	return products.Cantidad_disponible
}

//se creo la funcion NotRepeat
func NotRepeat(Struct CreateCompraInput) []DetalleInput {

	var auxiliar2 []DetalleInput
	var auxiliar3 DetalleInput
	var positionProducto int

	cont := 0
	for _, array := range Struct.DetalleInputs {

		flag := true
		for pos, array2 := range auxiliar2 {
			if array.Id_producto == array2.Id_producto {
				array2.Cantidad += array.Cantidad
				auxiliar3 = array2

				flag = false
				positionProducto = pos
			}
		}
		if flag == true {
			auxiliar3.Id_producto = array.Id_producto
			auxiliar3.Cantidad = array.Cantidad
			auxiliar2 = append(auxiliar2, auxiliar3)
		} else if flag == false {
			auxiliar2[positionProducto] = auxiliar3
		}
		flag = true
		cont++
	}

	return auxiliar2
}

func CreateCompra(c *gin.Context) {

	var cliente models.Cliente
	var producto models.Producto
	var input CreateCompraInput

	var CompraID int
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "error"})
		return
	}
	if err := models.DB.Where("id_cliente = ?", input.Id_cliente).First(&cliente).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	} else {
		values := NotRepeat(input)
		compra := models.Compra{Id_cliente: input.Id_cliente}
		models.DB.Create(&compra)
		CompraID = compra.Id_compra
		for _, array := range values {
			if err := models.DB.Where("id_producto = ?", array.Id_producto).First(&producto).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": array.Id_producto})
				return
			} else {
				returnFunctionUpdate := CantidadData(array.Id_producto)
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