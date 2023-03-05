package main

import (
	"tarea1/controllers"
	"tarea1/models"

	"github.com/gin-gonic/gin"
)

// probando
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
		grp1.GET("estadisticas", controllers.GetStats)
	}

	r.POST("api/clientes/iniciar_sesion", controllers.LoginClient)

	r.POST("api/compras", controllers.CreateCompra)

	r.Run(":5000")
}
