package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	service service.PaymentService
}

func NewPaymentController(service service.PaymentService) *PaymentController {
	return &PaymentController{service}
}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	var req entity.CreatePaymentReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := c.service.CreatePayment(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, payment)
}

func (c *PaymentController) GetPaymentByID(ctx *gin.Context) {
	id := ctx.Param("id")
	paymentID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	payment, err := c.service.GetPaymentByID(paymentID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

func (c *PaymentController) UpdatePayment(ctx *gin.Context) {
	var req entity.UpdatePaymentReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := c.service.UpdatePayment(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

func (c *PaymentController) DeletePayment(ctx *gin.Context) {
	id := ctx.Param("id")
	paymentID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	err = c.service.DeletePayment(paymentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}

func (c *PaymentController) GetAllPayments(ctx *gin.Context) {
	payments, err := c.service.GetAllPayments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}

func (c *PaymentController) UpdatePaymentStatus(ctx *gin.Context) {
	paymentID := ctx.Param("id")
	var req struct {
		Status string `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.service.UpdatePaymentStatus(paymentID, req.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment status updated successfully"})
}

func (c *PaymentController) GetPaymentReport(ctx *gin.Context) {
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

	// Panggil service untuk mendapatkan laporan pembayaran
	report, err := c.service.GetPaymentReport(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response JSON
	ctx.JSON(http.StatusOK, report)
}
