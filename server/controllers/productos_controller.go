package controllers

import (
	"log"
	"net/http"

	"tarea1/models"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Nombre              string `json:"nombre" binding:"required"`
	Cantidad_disponible int    `json:"cantidad_disponible" binding:"required"`
	Precio_unitario     int    `json:"precio_unitario" binding:"required"`
}

// Cuando actualizamos productos
type UpdateProductInput struct {
	Nombre              string `json:"nombre" binding:"required"`
	Cantidad_disponible int    `json:"cantidad_disponible" binding:"required"`
	Precio_unitario     int    `json:"precio_unitario" binding:"required"`
}

type Result struct {
	Id_producto int `json:"id_producto"`
}

func FindProducts(c *gin.Context) {
	var productos []models.Producto
	models.DB.Find(&productos)
	c.JSON(http.StatusOK, gin.H{"data": productos})
}

func CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	//fmt.Println(input)
	producto := models.Producto{Nombre: input.Nombre, Cantidad_disponible: input.Cantidad_disponible, Precio_unitario: input.Precio_unitario} //
	models.DB.Create(&producto)
	c.JSON(http.StatusOK, gin.H{"data": producto})
}

func GetProductByIDD(c *gin.Context) {
	id := c.Params.ByName("id")
	var producto models.Producto
	if err := models.DB.Where("id_producto = ?", id).First(&producto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RecordNotFound"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": producto})
	}
}

func DeleteProduct(c *gin.Context) {

	var producto models.Producto
	id := c.Params.ByName("id")
	models.DB.Where("id_producto = ?", id).Delete(producto)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateProduct(c *gin.Context) {
	var producto models.Producto
	id_producto := c.Params.ByName("id")
	if err := models.DB.Where("id_producto = ?", id_producto).First(&producto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RecordNotFound"})
		return
	}

	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&producto).Where("id_producto = ?", id_producto).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": id_producto})
}

func GetStats(c *gin.Context) {
	var Resultado Result
	var Resultado2 Result
	var Resultado3 Result
	var Resultado4 Result
	rs := models.DB.Raw("select id_producto from detalle group by id_producto order by sum(cantidad) DESC LIMIT 1;").Scan(&Resultado)
	rs2 := models.DB.Raw("select id_producto from detalle group by id_producto order by sum(cantidad) ASC LIMIT 1").Scan(&Resultado2)
	rs3 := models.DB.Raw("select d.id_producto from detalle d inner join producto p on p.id_producto = d.id_producto group by d.id_producto order by sum(d.cantidad)*p.precio_unitario desc limit 1").Scan(&Resultado3)
	rs4 := models.DB.Raw("select d.id_producto from detalle d inner join producto p on p.id_producto = d.id_producto group by d.id_producto order by sum(d.cantidad)*p.precio_unitario asc limit 1").Scan(&Resultado4)
	if rs.Error != nil && rs2.Error != nil && rs3.Error != nil && rs4.Error != nil {
		log.Println(rs.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"producto_mas_vendido": Resultado.Id_producto, "producto_menos_vendido": Resultado2.Id_producto, "producto_mas_ganancia": Resultado3.Id_producto, "producto_menos_ganancia": Resultado4.Id_producto})
}
