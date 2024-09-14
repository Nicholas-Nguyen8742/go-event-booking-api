package main

import (
	"event-booking-api/api/routes"
	"event-booking-api/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
