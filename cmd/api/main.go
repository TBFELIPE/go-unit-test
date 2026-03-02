package main

import (
	"go_unit_test/internal/controller"
	"go_unit_test/internal/repository"
	"go_unit_test/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	repo := &repository.RepositoryUser{}
	service := service.NewUserService(repo)
	handler := controller.NewUserHandler(service)

	router.POST("/get-user", handler.GetUserHandler)

	router.Run(":8080")
}
