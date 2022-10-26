package controllers

import (
	"api-alura/database"
	"api-alura/models"

	"github.com/gin-gonic/gin"
)

func Teste(c *gin.Context) {
	c.JSON(200, gin.H{
		"teste": "Anderson fazendo teste",
	})
}

func ExibeTodosVideos(c *gin.Context) {
	var videos []models.Video
	database.DB.Find(&videos)

	c.JSON(200, videos)
}
