package controller

import (
	"net/http"
	"strconv"
	"time"

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
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	bookings, err := c.service.GetAllBookings(limit, offset)
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

func (c *BookingController) GetBookingReport(ctx *gin.Context) {
	// Ambil parameter tanggal dari query string
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	var startDate, endDate time.Time
	var err error

	// Parse tanggal jika parameter diberikan
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
			return
		}
	}

	// Panggil service untuk mendapatkan laporan booking
	report, err := c.service.GetBookingReport(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response JSON
	ctx.JSON(http.StatusOK, report)
}

func (c *BookingController) GetAvailableDates(ctx *gin.Context) {
	// Ambil service_id dari query parameter
	serviceID, err := strconv.Atoi(ctx.Query("service_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service_id"})
		return
	}

	// Ambil tahun dan bulan dari query parameter
	year, err := strconv.Atoi(ctx.Query("year"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}

	month, err := strconv.Atoi(ctx.Query("month"))
	if err != nil || month < 1 || month > 12 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month"})
		return
	}

	// Panggil service untuk mendapatkan tanggal yang tersedia
	availableDates, err := c.service.GetAvailableDates(serviceID, year, month)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format tanggal ke dalam bentuk string (YYYY-MM-DD)
	var availableDatesStr []string
	for _, date := range availableDates {
		availableDatesStr = append(availableDatesStr, date.Format("2006-01-02"))
	}

	// Kembalikan response JSON
	ctx.JSON(http.StatusOK, gin.H{
		"service_id":      serviceID,
		"year":            year,
		"month":           month,
		"available_dates": availableDatesStr,
	})
}

func (c *BookingController) GetConfirmedBookingsForTechnician(ctx *gin.Context) {
	// Ambil user ID dari JWT token (asumsi sudah ada middleware yang menambahkan user ID ke context)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Konversi userID ke integer
	technicianID, ok := userID.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Panggil service untuk mendapatkan booking dengan status "Confirmed" untuk technician
	bookings, err := c.service.GetConfirmedBookingsForTechnician(technicianID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response JSON
	ctx.JSON(http.StatusOK, bookings)
}
