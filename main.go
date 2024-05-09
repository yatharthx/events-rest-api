package main

import (
	"github.com/yatharthx/events-rest-api/db"
	"github.com/yatharthx/events-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main()  {
  db.InitDB()
  server := gin.Default()
  routes.RegisterRoutes(server)

  server.Run(":8080")
}

