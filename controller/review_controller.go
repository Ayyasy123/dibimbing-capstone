package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	service service.ReviewService
}

func NewReviewController(service service.ReviewService) *ReviewController {
	return &ReviewController{service}
}

func (c *ReviewController) CreateReview(ctx *gin.Context) {
	var req entity.CreateReviewReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := c.service.CreateReview(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, review)
}

func (c *ReviewController) GetReviewByID(ctx *gin.Context) {
	id := ctx.Param("id")
	reviewID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	review, err := c.service.GetReviewByID(reviewID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	ctx.JSON(http.StatusOK, review)
}

func (c *ReviewController) UpdateReview(ctx *gin.Context) {
	var req entity.UpdateReviewReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := c.service.UpdateReview(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, review)
}

func (c *ReviewController) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")
	reviewID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	err = c.service.DeleteReview(reviewID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

func (c *ReviewController) GetAllReviews(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	reviews, err := c.service.GetAllReviews(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}

func (c *ReviewController) GetReviewReport(ctx *gin.Context) {
	// Ambil parameter tanggal dari query string
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	// Ambil parameter service_id dari query string
	serviceIDStr := ctx.Query("service_id")
	serviceID, _ := strconv.Atoi(serviceIDStr) // Jika tidak ada, serviceID akan 0

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

	// Panggil service untuk mendapatkan laporan review
	report, err := c.service.GetReviewReport(startDate, endDate, serviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response JSON
	ctx.JSON(http.StatusOK, report)
}
