package controllers

import (
	"api-alura/database"
	"api-alura/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodasCategorias(c *gin.Context) {
	var categorias []models.Categoria
	database.DB.Find(&categorias)

	c.JSON(http.StatusOK, categorias)
}

func ExibeUmaCaterogia(c *gin.Context) {
	var categoria models.Categoria
	id := c.Params.ByName("id")
	database.DB.First(&categoria, id)

	if categoria.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"404": "Categoria não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, categoria)
}

func CriarCateria(c *gin.Context) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&categoria)
	c.JSON(http.StatusOK, categoria)
}

func AlterarCategoria(c *gin.Context) {
	var categoria models.Categoria
	id := c.Params.ByName("id")

	database.DB.First(&categoria, id)

	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&categoria).UpdateColumns(categoria)
	c.JSON(http.StatusOK, categoria)
}

func DeletaCategoria(c *gin.Context) {
	var categorias []models.Video

	id := c.Params.ByName("id")

	database.DB.Delete(&categorias, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Categoria deletado com sucesso",
	})
}
