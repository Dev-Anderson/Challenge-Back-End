package main

import (
	"api-alura/database"
	"fmt"
)

func main() {
	fmt.Println("Iniciando projeto do back-end Alura")
	database.GetDatabase()
}
