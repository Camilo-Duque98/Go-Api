package controllers

import (
	"fmt"
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
	var Resultado Result
	var Resultado2 Result
	var Resultado3 Result
	var Resultado4 Result

	//db, err := gorm.Open("mysql", "root:mysql1234567890@tcp(localhost:3306)/tarea_1_sd")
	//if err != nil {
	//	panic("Failed to connect to database!")
	//}
	//rs := models.DB.Raw("SELECT id_producto FROM producto WHERE id_producto = 8").Scan(&Resultado)
	rs := models.DB.Raw("SELECT id_producto FROM detalle GROUP BY id_producto ORDER BY COUNT(id_producto) DESC LIMIT 1").Scan(&Resultado)
	rs2 := models.DB.Raw("SELECT id_producto FROM detalle GROUP BY id_producto ORDER BY COUNT(id_producto) ASC LIMIT 1").Scan(&Resultado2)
	rs3 := models.DB.Raw("select p.nombre,d.id_producto, sum(d.cantidad)*p.precio_unitario as total from detalle d inner join producto p on p.id_producto = d.id_producto group by d.id_producto order by count(d.id_producto) desc limit 1;").Scan(&Resultado3)
	rs4 := models.DB.Raw("select p.nombre,d.id_producto, sum(d.cantidad)*p.precio_unitario as total from detalle d inner join producto p on p.id_producto = d.id_producto group by d.id_producto order by count(d.id_producto) asc limit 1;").Scan(&Resultado4)
	if rs.Error != nil && rs2.Error != nil && rs3.Error != nil && rs4.Error != nil {
		log.Println(rs.Error)
		return
	}
	/*if rs2.Error != nil && {
		log.Println(rs2.Error)
		return
	}
	if rs3.Error != nil {
		log.Println(rs3.Error)
		return
	}
	if rs4.Error != nil {
		log.Println(rs4.Error)
		return
	}*/
	fmt.Println(Resultado)
	c.JSON(http.StatusOK, gin.H{"producto_mas_vendido": Resultado.Id_producto, "producto_menos_vendido": Resultado2.Id_producto, "producto_mas_ganancia": Resultado3.Id_producto, "producto_menos_ganancia": Resultado4.Id_producto})

}
