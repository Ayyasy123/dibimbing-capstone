package controller

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"github.com/Ayyasy123/dibimbing-capstone.git/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var req entity.RegisterUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRes, err := c.userService.Register(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, userRes)
}

func (c *UserController) Login(ctx *gin.Context) {
	var req entity.LoginUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRes, token, err := c.userService.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response dalam format yang diinginkan
	ctx.JSON(http.StatusOK, gin.H{
		"user":  userRes,
		"token": token,
	})
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	userRes, err := c.userService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, userRes)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	users, err := c.userService.GetAllUsers(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var req entity.UpdateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRes, err := c.userService.UpdateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userRes)
}

func (c *UserController) RegisterAsTechnician(ctx *gin.Context) {
	var req entity.RegisterAsTechnicianReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	technicianRes, err := c.userService.RegisterAsTechnician(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, technicianRes)
}

func (c *UserController) UpdateTechnician(ctx *gin.Context) {
	var req entity.UpdateTechnicianReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	technicianRes, err := c.userService.UpdateTechnician(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, technicianRes)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (c *UserController) RegisterAsAdmin(ctx *gin.Context) {
	var req entity.RegisterUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk register sebagai admin
	userRes, err := c.userService.RegisterAsAdmin(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response
	ctx.JSON(http.StatusCreated, userRes)
}

func (c *UserController) GetUserRoleReport(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	report, err := c.userService.GetUserRoleReport(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}
