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

//se creo una estructura para guardar supuestos elementos repetidos
type RepeatNotElement struct {
	Id_producto int
	cantidad    int
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

//se creo la funcion NotRepeat
func NotRepeat(Struct CreateCompraInput) []DetalleInput {

	//var auxiliar CreateCompraInput

	var auxiliar2 []DetalleInput
	var auxiliar3 DetalleInput
	//var contadorProducto int
	var positionProducto int

	cont := 0
	for _, array := range Struct.DetalleInputs { //{9 1}
		//fmt.Println("Printiando: ", array)

		flag := true
		for pos, array2 := range auxiliar2 { // {4 2} {9 1}
			if array.Id_producto == array2.Id_producto { //4 =9 -> hay que arreglar acá
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
		//aca estamos probando con printeos
		//fmt.Println("Probando detalles: ", input.DetalleInputs)
		values := NotRepeat(input)
		fmt.Println("Viendo si es que funciona", values)
		compra := models.Compra{Id_cliente: input.Id_cliente}
		models.DB.Create(&compra)
		CompraID = compra.Id_compra

		for _, array := range values {
			fmt.Println("sdjjfkdhasljfhdsakdfhakshjflkahjsfdkjdlas: ", array.Id_producto)
			fmt.Println("Cantidadfsdfsafdsfsafds: ", array.Cantidad)
			if err := models.DB.Where("id_producto = ?", array.Id_producto).First(&producto).Error; err != nil {
				//c.JSON(http.StatusBadRequest, gin.H{"error": "Producto no encontrado"})
				c.JSON(http.StatusBadRequest, gin.H{"error": array.Id_producto})
				return
			} else {
				//fmt.Println(array.Id_producto)
				returnFunctionUpdate := CantidadData(array.Id_producto)
				//fmt.Println(returnFunctionUpdate)
				if array.Cantidad <= returnFunctionUpdate {
					returnUpdateProductIncrease := UpdateProductIncrease(array.Id_producto, returnFunctionUpdate, array.Cantidad)
					fmt.Println(returnUpdateProductIncrease)
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "productos insuficientes"})
				}
				detalle := models.Detalle{Id_compra: CompraID, Id_producto: array.Id_producto, Cantidad: array.Cantidad}
				fmt.Println(detalle)
				models.DB.Create(&detalle)
			}

		}

		c.JSON(http.StatusOK, gin.H{"id_compra": compra.Id_compra})
	}

}
