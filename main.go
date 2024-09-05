package main

import (
	"alura/gin-api-rest/database"
	"alura/gin-api-rest/routes"
)

func main() {
	database.ConectaBanco()
	routes.HandleRequests()
}
