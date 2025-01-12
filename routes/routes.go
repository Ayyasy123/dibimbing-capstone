package routes

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/controller"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(db *gorm.DB, router *gin.Engine) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.GET("/users/:id", userController.GetUserByID)
	router.GET("/users", userController.GetAllUsers)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// userRoutes := router.Group("/users")
	// {
	// 	// userRoutes.POST("/register", userController.RegisterUser)
	// 	userRoutes.GET("/:id", userController.GetUserByID)
	// 	userRoutes.PUT("/:id", userController.UpdateUser)
	// 	userRoutes.DELETE("/:id", userController.DeleteUser)

	// 	// userRoutes.GET("/users/:id", userController.GetUserByID)
	// 	userRoutes.GET("/all", userController.GetAllUsers)
	// 	// userRoutes.POST("/users", userController.CreateUser)
	// 	// userRoutes.PUT("/users", userController.UpdateUser)
	// 	// userRoutes.DELETE("/users/:id", userController.DeleteUser)

	// 	userRoutes.POST("/register", userController.Register)
	// 	userRoutes.GET("/login", userController.Login)

	// }

}

func SetupServiceRoutes(db *gorm.DB, router *gin.Engine) {
	serviceRepo := repository.NewServiceRepository(db)
	serviceService := service.NewServiceService(serviceRepo)
	serviceController := controller.NewServiceController(serviceService)

	router.POST("/services", serviceController.CreateService)
	router.GET("/services/:id", serviceController.GetServiceByID)
	router.PUT("/services", serviceController.UpdateService)
	router.DELETE("/services/:id", serviceController.DeleteService)
	router.GET("/services", serviceController.GetAllServices)
}
