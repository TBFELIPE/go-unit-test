package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "go_unit_test/internal/dto"
	// "go_unit_test/internal/service"
)

type UserService interface {
	GetUserService(id string) (dto.User, error)
}

type UserHandlers struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandlers {
	return &UserHandlers{service: s}
}

func (c *UserHandlers) GetUserHandler(ctx *gin.Context) {

	var req dto.Request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error while binding request json": err.Error()})
		return
	}

	id := req.Id

	user, err := c.service.GetUserService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
