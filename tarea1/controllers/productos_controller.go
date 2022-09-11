package controllers

import (
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
	"tarea1/models"
)
type CreateProductInput struct {
	Nombre  string `json:"nombre" binding:"required"`
	Cantidad_disponible int `json:"cantidad_disponible" binding:"required"`
	Precio_unitario int `json:"precio_unitario" binding:"required"`
}

type UpdateProductInput struct{
	Nombre  string `json:"nombre" binding:"required"`
	Cantidad_disponible int `json:"cantidad_disponible" binding:"required"`
	Precio_unitario int `json:"precio_unitario" binding:"required"`
}

func FindProducts(c *gin.Context) {
	var productos []models.Producto
	models.DB.Find(&productos)
	c.JSON(http.StatusOK, gin.H{"data": productos})
}

func CreateProduct(c *gin.Context){
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	//Create cliente
	producto := models.Producto{Nombre: input.Nombre, Cantidad_disponible: input.Cantidad_disponible, Precio_unitario: input.Precio_unitario} //
	models.DB.Create(&producto)
	//fmt.Println(input)
	c.JSON(http.StatusOK, gin.H{"data": producto})
}

//-------------------------------------------------------------------------
/*func GetProductByID(producto *Producto, id_producto string) (err error) {
	if err = models.DB.Where("id_producto = ?", id_producto).First(producto).Error; err != nil {
		return err
	}
	return nil
}*/

func GetProductByIDD(c *gin.Context) {
	id := c.Params.ByName("id")
	var producto models.Producto
	//err := models.GetProductByID(&producto, id_producto)
	if err := models.DB.Where("id_producto = ?", id).First(&producto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"RecordNotFound"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": producto})
	}
}

func DeleteProduct(c *gin.Context){
	var producto models.Producto
	id := c.Params.ByName("id")
	//err := models.DeleteUser(&user, id)
	//Config.DB.Where("id = ?", id).Delete(user)
	//return nil
	err := models.DB.Where("id_producto = ?",id).Delete(producto)
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

//--------------------
/*func FindProduct(c *gin.Context) {  // Get model if exist
	var producto models.Producto
	id_producto := c.Request.URL.Query().Get("id_producto")
	//id_producto := c.Params.ByName("id_producto")
	if err := models.DB.Where("id_producto = ?", id_producto).First(&producto).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": producto})
  
  }*/

func UpdateProduct(c *gin.Context) {
	var producto models.Producto
	id_producto := c.Request.URL.Query().Get("id_producto")
	if err := models.DB.Where("id_producto = ?", id_producto).First(&producto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"RecordNotFound"})
		return
	}

	//Validamos el input

	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	models.DB.Model(&producto).Updates(input)
	
	c.JSON(http.StatusOK, gin.H{"data": producto})
}