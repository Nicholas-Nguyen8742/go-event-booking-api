package main

import (
	"event-booking-api/api/routes"
	"event-booking-api/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.App()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
