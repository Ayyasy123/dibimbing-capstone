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

	// r.POST("/register", service.CreateUserHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.SetupUserRoutes(config.DB, r)

	// r.GET("/users", userController.GetAllUsers)
	// r.GET("/users/:id", userController.GetUserByID)
	// r.POST("/users", userController.CreateUser)
	// r.PUT("/users", userController.UpdateUser)
	// r.DELETE("/users/:id", userController.DeleteUser)

	// r.POST("/register", userController.Register)
	// r.GET("/login", userController.Login)

	// Start the Server
	log.Println("Server is running on http://localhost:8080")
	//Start server
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
