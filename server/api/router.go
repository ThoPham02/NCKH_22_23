package api

import (
	"github/ThoPham02/research_management/api/handler"
	"github/ThoPham02/research_management/api/middleware"
	"github/ThoPham02/research_management/api/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine, svc *service.ServiceContext) {
	router.Use(middleware.CorsMiddleware())
	// router.Use(middleware.)
	// user api
	router.POST("api/user/login", handler.UserLoginHandler(svc))
	router.POST("api/user/register", handler.UserRegisterHandler(svc))
	router.GET("api/user/info", middleware.UserAuthentication(svc), handler.GetUserInfoHandler(svc))

	//topic api
}
