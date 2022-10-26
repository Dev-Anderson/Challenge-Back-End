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

type env struct {
	host     string
	port     string
	user     string
	pass     string
	database string
}

func GetDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var e = env{}

	e.host = os.Getenv("HOST")
	e.port = os.Getenv("PORT")
	e.user = os.Getenv("USER")
	e.pass = os.Getenv("PASS")
	e.database = os.Getenv("DATABASE")

	conexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", e.host, e.user, e.pass, e.database, e.port)
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Video{})
}
