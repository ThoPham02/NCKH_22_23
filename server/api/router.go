package api

import (
	"github/ThoPham02/research_management/api/handler"
	"github/ThoPham02/research_management/api/middleware"
	"github/ThoPham02/research_management/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine, svc *service.ServiceContext) {
	router.Use(middleware.CorsMiddleware())
	router.POST("user/login", handler.UserRegister(svc))
	router.GET("/user", middleware.AdminAuthentication(svc), handler.GetUserHandler(svc))
}
