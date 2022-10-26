package routes

import (
	"api-alura/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/", controllers.Teste)
	r.GET("/videos", controllers.ExibeTodosVideos)
	r.GET("/videos/:id", controllers.ExibeUmVideo)
	r.Run()
}
