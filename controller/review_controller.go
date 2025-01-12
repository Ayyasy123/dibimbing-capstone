package controller

import (
	"net/http"
	"strconv"

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
	reviews, err := c.service.GetAllReviews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}
