package controller

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	service service.BookingService
}

func NewBookingController(service service.BookingService) *BookingController {
	return &BookingController{service}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var req entity.CreateBookingReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking, err := c.service.CreateBooking(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, booking)
}

func (c *BookingController) GetBookingByID(ctx *gin.Context) {
	id := ctx.Param("id")
	bookingID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	booking, err := c.service.GetBookingByID(bookingID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (c *BookingController) UpdateBooking(ctx *gin.Context) {
	var req entity.UpdateBookingReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking, err := c.service.UpdateBooking(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (c *BookingController) DeleteBooking(ctx *gin.Context) {
	id := ctx.Param("id")
	bookingID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	err = c.service.DeleteBooking(bookingID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}

func (c *BookingController) GetAllBookings(ctx *gin.Context) {
	bookings, err := c.service.GetAllBookings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bookings)
}

func (c *BookingController) GetBookingsByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	bookings, err := c.service.GetBookingsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bookings)
}

func (c *BookingController) GetBookingsByServiceID(ctx *gin.Context) {
	serviceID, err := strconv.Atoi(ctx.Param("service_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	bookings, err := c.service.GetBookingsByServiceID(serviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bookings)
}

func (c *BookingController) UpdateBookingStatus(ctx *gin.Context) {
	bookingID := ctx.Param("id")
	var req struct {
		Status string `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.service.UpdateBookingStatus(bookingID, req.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Booking status updated successfully"})
}
