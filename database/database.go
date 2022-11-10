package database

import (
	"api-alura/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func GetDatabase() {
	err := godotenv.Load()
	e := models.Env{}
	if err != nil {
		panic(err)
	}

	e.Host = os.Getenv("HOST")
	e.User = os.Getenv("USER")
	e.Pass = os.Getenv("PASS")
	e.Database = os.Getenv("DATABASE")

	conexao := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", e.Host, e.User, e.Pass, e.Database)
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Video{}, &models.Categoria{})
}
