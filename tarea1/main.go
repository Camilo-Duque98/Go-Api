package main

import (
  "github.com/gin-gonic/gin"
  "tarea1/models"
  "tarea1/controllers"
)

func main() {
  r := gin.Default()
  models.ConnectDatabase()

  grp1 := r.Group("/api")
	{
		grp1.GET("productos", controllers.FindProducts)
		grp1.POST("producto", controllers.CreateProduct)
		grp1.GET("producto/:id", controllers.GetProductByIDD)
		grp1.PUT("producto/:id", controllers.UpdateProduct)
		grp1.DELETE("producto/:id", controllers.DeleteProduct)
	}

  // Productos
  /*r.GET("/api/productos", controllers.FindProducts)
  r.POST("/api/producto", controllers.CreateProduct)
  r.GET("/api/producto/:id_producto", controllers.FindProduct) 
  r.PUT("/api/producto/:id_producto", controllers.UpdateProduct) 
*/
  // Clientes

  r.POST("api/clientes/iniciar_sesion", controllers.LoginClient)
  r.Run()

  //
}