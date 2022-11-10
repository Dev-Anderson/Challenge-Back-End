package routes

import (
	"api-alura/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		videos := main.Group("videos")
		{
			videos.GET("/", controllers.ExibeTodosVideos)
			videos.GET("/:id", controllers.ExibeUmVideo)
			videos.POST("/", controllers.CriarVideo)
			videos.PATCH("/:id", controllers.AlterarUmVideo)
			videos.DELETE("/:id", controllers.DeletaVideo)
		}
		categorias := main.Group("categorias")
		{
			categorias.GET("/", controllers.ExibeTodasCategorias)
			categorias.GET("/:id", controllers.ExibeUmaCaterogia)
			categorias.POST("/", controllers.CriarCateria)
			categorias.PATCH("/:id", controllers.AlterarCategoria)
			categorias.DELETE("/:id", controllers.DeletaCategoria)
		}
	}

	return router
}
