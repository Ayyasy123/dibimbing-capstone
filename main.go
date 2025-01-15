package main

import (
	"log"
	"net/http"

	"github.com/Ayyasy123/dibimbing-capstone.git/config"
	"github.com/Ayyasy123/dibimbing-capstone.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Setup Gin Router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.SetupUserRoutes(config.DB, r)
	routes.SetupServiceRoutes(config.DB, r)
	routes.SetupBookingRoutes(config.DB, r)
	routes.SetupPaymentRoutes(config.DB, r)
	routes.SetupReviewRoutes(config.DB, r)

	// Start the Server
	log.Println("Server is running on http://localhost:8080")
	//Start server
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
