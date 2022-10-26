package controllers

import (
	"api-alura/database"
	"api-alura/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Teste(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"teste": "Anderson fazendo teste",
	})
}

func ExibeTodosVideos(c *gin.Context) {
	var videos []models.Video
	database.DB.Find(&videos)

	c.JSON(http.StatusOK, videos)
}

func ExibeUmVideo(c *gin.Context) {
	var video models.Video
	id := c.Params.ByName("id")
	database.DB.First(&video, id)

	if video.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"404": "Vídeo não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, video)
}

func CriarVideo(c *gin.Context) {
	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&video)
	c.JSON(http.StatusOK, video)
}
