package routes

import (
	"github.com/Ayyasy123/dibimbing-capstone.git/controller"
	"github.com/Ayyasy123/dibimbing-capstone.git/middleware"
	"github.com/Ayyasy123/dibimbing-capstone.git/repository"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(db *gorm.DB, router *gin.Engine) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Public routes (no authentication required)
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	// Endpoint untuk register sebagai admin (hanya bisa diakses oleh admin)
	router.POST("/register-admin", userController.RegisterAsAdmin)

	// Protected routes (require JWT authentication)
	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.JWTAuth())
	{
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.GET("", userController.GetAllUsers)
		userRoutes.PUT("", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)

		// Role-based routes
		userRoutes.POST("/register-technician", userController.RegisterAsTechnician)
		userRoutes.PUT("/update-technician", middleware.RoleAuth("technician", "admin"), userController.UpdateTechnician)

	}
}

func SetupServiceRoutes(db *gorm.DB, router *gin.Engine) {
	serviceRepo := repository.NewServiceRepository(db)
	serviceService := service.NewServiceService(serviceRepo)
	serviceController := controller.NewServiceController(serviceService)

	// Protected routes (require JWT authentication)
	serviceRoutes := router.Group("/services")
	serviceRoutes.Use(middleware.JWTAuth())
	{
		serviceRoutes.POST("", middleware.RoleAuth("technician"), serviceController.CreateService)
		serviceRoutes.GET("/:id", serviceController.GetServiceByID)
		serviceRoutes.PUT("", middleware.RoleAuth("technician"), serviceController.UpdateService)
		serviceRoutes.DELETE("/:id", middleware.RoleAuth("technician"), serviceController.DeleteService)
		serviceRoutes.GET("", serviceController.GetAllServices)
		serviceRoutes.GET("/user/:user_id", serviceController.GetServicesByUserID)
		serviceRoutes.GET("/search", serviceController.SearchServices) /// services/search?search=plumbing&min_price=10000&max_price=50000
	}
}

func SetupBookingRoutes(db *gorm.DB, router *gin.Engine) {
	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo)
	bookingController := controller.NewBookingController(bookingService)

	// Protected routes (require JWT authentication)
	bookingRoutes := router.Group("/bookings")
	bookingRoutes.Use(middleware.JWTAuth())
	{
		bookingRoutes.GET("", bookingController.GetAllBookings)
		bookingRoutes.GET("/:id", bookingController.GetBookingByID)
		bookingRoutes.POST("", bookingController.CreateBooking)
		bookingRoutes.PUT("", bookingController.UpdateBooking)
		bookingRoutes.DELETE("/:id", bookingController.DeleteBooking)
		bookingRoutes.GET("/user/:user_id", bookingController.GetBookingsByUserID)
		bookingRoutes.GET("/service/:service_id", bookingController.GetBookingsByServiceID)
		bookingRoutes.PUT("/:id/status", bookingController.UpdateBookingStatus)
		bookingRoutes.GET("/reports", bookingController.GetBookingReport)
	}
}

func SetupPaymentRoutes(db *gorm.DB, router *gin.Engine) {
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentController := controller.NewPaymentController(paymentService)

	// Protected routes (require JWT authentication)
	paymentRoutes := router.Group("/payments")
	paymentRoutes.Use(middleware.JWTAuth())
	{
		paymentRoutes.GET("", paymentController.GetAllPayments)
		paymentRoutes.GET("/:id", paymentController.GetPaymentByID)
		paymentRoutes.POST("", paymentController.CreatePayment)
		paymentRoutes.PUT("", paymentController.UpdatePayment)
		paymentRoutes.DELETE("/:id", paymentController.DeletePayment)
		paymentRoutes.PUT("/:id/status", paymentController.UpdatePaymentStatus)
		paymentRoutes.GET("/reports", paymentController.GetPaymentReport)
	}
}

func SetupReviewRoutes(db *gorm.DB, router *gin.Engine) {
	reviewRepo := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo)
	reviewController := controller.NewReviewController(reviewService)

	// Protected routes (require JWT authentication)
	reviewRoutes := router.Group("/reviews")
	reviewRoutes.Use(middleware.JWTAuth())
	{
		reviewRoutes.GET("", reviewController.GetAllReviews)
		reviewRoutes.GET("/:id", reviewController.GetReviewByID)
		reviewRoutes.POST("", reviewController.CreateReview)
		reviewRoutes.PUT("", reviewController.UpdateReview)
		reviewRoutes.DELETE("/:id", reviewController.DeleteReview)
	}
}
