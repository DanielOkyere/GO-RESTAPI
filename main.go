package main

import (
	"danielokyere/RESTCRUD/db"
	"danielokyere/RESTCRUD/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
