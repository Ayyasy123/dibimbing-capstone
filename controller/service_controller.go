package controller

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
)

// type ServiceController interface {
// 	CreateService(ctx *gin.Context)
// 	GetServiceByID(ctx *gin.Context)
// 	UpdateService(ctx *gin.Context)
// 	DeleteService(ctx *gin.Context)
// 	GetAllServices(ctx *gin.Context)
// 	GetServicesByUserID(ctx *gin.Context)
// 	SearchServices(ctx *gin.Context)
// }

type ServiceController struct {
	serviceService service.ServiceService
}

func NewServiceController(serviceService service.ServiceService) *ServiceController {
	return &ServiceController{serviceService: serviceService}
}

func (c *ServiceController) CreateService(ctx *gin.Context) {
	var req entity.CreateServiceReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := c.serviceService.CreateService(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	ctx.JSON(http.StatusCreated, serviceRes)
}

func (c *ServiceController) GetServiceByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	service, err := c.serviceService.GetServiceByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, service)
}

func (c *ServiceController) UpdateService(ctx *gin.Context) {
	var req entity.UpdateServiceReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := c.serviceService.UpdateService(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	ctx.JSON(http.StatusOK, serviceRes)
}

func (c *ServiceController) DeleteService(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = c.serviceService.DeleteService(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

func (c *ServiceController) GetAllServices(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	services, err := c.serviceService.GetAllServices(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}

func (c *ServiceController) GetServicesByUserID(ctx *gin.Context) {
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

func (c *ServiceController) SearchServices(ctx *gin.Context) {
	searchQuery := ctx.Query("search") // Ambil parameter query string "search"
	minPriceStr := ctx.Query("min_price")
	maxPriceStr := ctx.Query("max_price")

	var minPrice, maxPrice int
	var err error

	// Jika min_price tidak disebutkan, set ke 0
	if minPriceStr == "" {
		minPrice = 0
	} else {
		minPrice, err = strconv.Atoi(minPriceStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_price"})
			return
		}
	}

	// Jika max_price tidak disebutkan, set ke 100 juta
	if maxPriceStr == "" {
		maxPrice = 100000000
	} else {
		maxPrice, err = strconv.Atoi(maxPriceStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid max_price"})
			return
		}
	}

	if minPrice > maxPrice {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "min_price must be less than or equal to max_price"})
		return
	}

	services, err := c.serviceService.SearchServices(searchQuery, minPrice, maxPrice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}

func (c *ServiceController) GetServiceCostReport(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	report, err := c.serviceService.GetServiceCostReport(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}
