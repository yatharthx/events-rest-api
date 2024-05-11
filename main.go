package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yatharthx/events-rest-api/db"
	"github.com/yatharthx/events-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
