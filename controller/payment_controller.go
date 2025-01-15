package controller

import (
	"net/http"
	"strconv"

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
