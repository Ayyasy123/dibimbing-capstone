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

func SetupBookingRoutes(db *gorm.DB, router *gin.Engine) {
	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo)
	bookingController := controller.NewBookingController(bookingService)

	// Booking routes
	router.GET("/bookings", bookingController.GetAllBookings)
	router.GET("/bookings/:id", bookingController.GetBookingByID)
	router.POST("/bookings", bookingController.CreateBooking)
	router.PUT("/bookings", bookingController.UpdateBooking)
	router.DELETE("/bookings/:id", bookingController.DeleteBooking)
}

func SetupPaymentRoutes(db *gorm.DB, router *gin.Engine) {
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentController := controller.NewPaymentController(paymentService)

	// Payment routes
	router.GET("/payments", paymentController.GetAllPayments)
	router.GET("/payments/:id", paymentController.GetPaymentByID)
	router.POST("/payments", paymentController.CreatePayment)
	router.PUT("/payments", paymentController.UpdatePayment)
	router.DELETE("/payments/:id", paymentController.DeletePayment)
}

func SetupReviewRoutes(db *gorm.DB, router *gin.Engine) {
	reviewRepo := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo)
	reviewController := controller.NewReviewController(reviewService)

	// Review routes
	router.GET("/reviews", reviewController.GetAllReviews)
	router.GET("/reviews/:id", reviewController.GetReviewByID)
	router.POST("/reviews", reviewController.CreateReview)
	router.PUT("/reviews", reviewController.UpdateReview)
	router.DELETE("/reviews/:id", reviewController.DeleteReview)
}
