package controller

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
)

type ServiceController interface {
	CreateService(c *gin.Context)
	GetServiceByID(c *gin.Context)
	UpdateService(c *gin.Context)
	DeleteService(c *gin.Context)
	GetAllServices(c *gin.Context)
	GetServicesByUserID(ctx *gin.Context)
	SearchServices(c *gin.Context)
}

type serviceController struct {
	serviceService service.ServiceService
}

func NewServiceController(serviceService service.ServiceService) ServiceController {
	return &serviceController{serviceService}
}

func (ctrl *serviceController) CreateService(c *gin.Context) {
	var req entity.CreateServiceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := ctrl.serviceService.CreateService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Gunakan ServiceRes untuk menghilangkan User dari response
	serviceRes := entity.ServiceRes{
		ID:          service.ID,
		UserID:      service.UserID,
		Name:        service.Name,
		Description: service.Description,
		Cost:        service.Cost,
		CreatedAt:   service.CreatedAt,
		UpdatedAt:   service.UpdatedAt,
	}

	c.JSON(http.StatusCreated, serviceRes)
}

func (ctrl *serviceController) GetServiceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	service, err := ctrl.serviceService.GetServiceByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

func (ctrl *serviceController) UpdateService(c *gin.Context) {
	var req entity.UpdateServiceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := ctrl.serviceService.UpdateService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Gunakan ServiceRes untuk menghilangkan User dari response
	serviceRes := entity.ServiceRes{
		ID:          service.ID,
		UserID:      service.UserID,
		Name:        service.Name,
		Description: service.Description,
		Cost:        service.Cost,
		CreatedAt:   service.CreatedAt,
		UpdatedAt:   service.UpdatedAt,
	}

	c.JSON(http.StatusOK, serviceRes)
}

func (ctrl *serviceController) DeleteService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = ctrl.serviceService.DeleteService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

func (ctrl *serviceController) GetAllServices(c *gin.Context) {
	services, err := ctrl.serviceService.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}

func (c *serviceController) GetServicesByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	services, err := c.serviceService.GetServicesByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}

func (ctrl *serviceController) SearchServices(c *gin.Context) {
	searchQuery := c.Query("search") // Ambil parameter query string "search"
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")

	var minPrice, maxPrice int
	var err error

	// Jika min_price tidak disebutkan, set ke 0
	if minPriceStr == "" {
		minPrice = 0
	} else {
		minPrice, err = strconv.Atoi(minPriceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_price"})
			return
		}
	}

	// Jika max_price tidak disebutkan, set ke 100 juta
	if maxPriceStr == "" {
		maxPrice = 100000000
	} else {
		maxPrice, err = strconv.Atoi(maxPriceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid max_price"})
			return
		}
	}

	if minPrice > maxPrice {
		c.JSON(http.StatusBadRequest, gin.H{"error": "min_price must be less than or equal to max_price"})
		return
	}

	services, err := ctrl.serviceService.SearchServices(searchQuery, minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}
