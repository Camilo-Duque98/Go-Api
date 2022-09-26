package controllers

import (
	"log"
	"net/http"

	"tarea1/models"

	"github.com/gin-gonic/gin"
)

// Estructura de los productos
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

	//Validamos el input

	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&producto).Where("id_producto = ?", id_producto).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": id_producto})
}

// funcion de estadistica
func GetStats(c *gin.Context) {
	var id_producto int64
	//db, err := gorm.Open("mysql", "root:mysql1234567890@tcp(localhost:3306)/tarea_1_sd")
	//if err != nil {
	//	panic("Failed to connect to database!")
	//}
	rs := models.DB.Raw("select id_producto from detalle group by id_producto order by count(id_producto) desc limit 1").Scan(&id_producto)
	if rs.Error != nil {
		log.Println(rs.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id_producto})

}
