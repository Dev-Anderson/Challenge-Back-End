package main

import (
	"api-alura/database"
	"api-alura/server"
)

func main() {
	database.GetDatabase()

	s := server.NewServer()
	s.Run()
}
