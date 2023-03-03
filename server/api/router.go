package api

import (
	"github/ThoPham02/research_management/api/handler"
	"github/ThoPham02/research_management/api/middleware"
	"github/ThoPham02/research_management/api/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine, svc *service.ServiceContext) {
	router.Use(middleware.CorsMiddleware())
	router.POST("user/login", handler.UserRegister(svc))
	router.GET("/user", middleware.AdminAuthentication(svc), handler.GetUserHandler(svc))

	router.GET("/topics", handler.GetListTopicsHandler(svc))
	router.GET("/departments", handler.GetStudentHandler(svc))
	router.GET("/faculties", handler.GetUserHandler(svc))
	router.GET("/students", handler.GetUserHandler(svc))
	router.GET("/lectures", handler.GetUserHandler(svc))

	router.GET("/department/info", middleware.UserAuthentication(svc), handler.GetStudentHandler(svc))
	router.GET("/faculty/info", middleware.UserAuthentication(svc), handler.GetUserHandler(svc))
	router.GET("/student/info", middleware.UserAuthentication(svc), handler.GetUserHandler(svc))
	router.GET("/lecture/info", middleware.UserAuthentication(svc), handler.GetUserHandler(svc))
}
