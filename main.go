package main

import (
	"api-alura/database"
	"api-alura/routes"
)

func main() {
	database.GetDatabase()

	routes.HandleRequest()
}
