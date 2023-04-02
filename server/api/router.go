package api

import (
	"github/ThoPham02/research_management/api/middleware"
	"github/ThoPham02/research_management/api/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine, svc *service.ServiceContext) {
	router.Use(middleware.CorsMiddleware())
	// router.POST("user/login", handler.UserRegister(svc))
	// router.GET("/user", middleware.AdminAuthentication(svc), handler.GetUserHandler(svc))
	// // router.GET("/user/changepassword", handler.ChangePassword(svc)) // change all password to hash password

	// router.GET("/topics", handler.GetListTopicsHandler(svc))
	// router.GET("/departments", handler.GetListDepartmentsHandler(svc))
	// router.GET("/faculties", handler.GetListFacultiesHandler(svc))

	// router.GET("/department/info", middleware.UserAuthentication(svc), handler.GetDepartmentInfoHandler(svc))
	// router.GET("/faculty/info", middleware.UserAuthentication(svc), handler.GetFacultyInfoHandler(svc))
	// router.GET("/student/info", middleware.UserAuthentication(svc), handler.GetStudentInfoHandler(svc))
	// router.GET("/lecture/info", middleware.UserAuthentication(svc), handler.GetLecturerInfoHandler(svc))
}
