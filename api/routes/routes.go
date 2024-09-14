package routes

import (
	"event-booking-api/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	protectedRoutes := server.Group("/")
	protectedRoutes.Use(middleware.Authenticate)
	protectedRoutes.POST("/events", createEvent)
	protectedRoutes.PUT("/events/:id", updateEvent)
	protectedRoutes.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
