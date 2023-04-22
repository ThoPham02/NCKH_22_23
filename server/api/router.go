package api

import (
	"github/ThoPham02/research_management/api/handler"
	"github/ThoPham02/research_management/api/middleware"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(router *gin.Engine, svc *service.ServiceContext) {
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.CorsMiddleware())

	// user api
	router.POST("api/user/login", handler.UserLoginHandler(svc))                                        // đăng nhập // done
	router.POST("api/user/register", handler.UserRegisterHandler(svc))                                  // đăng kí // done
	router.GET("api/user/info", middleware.UserAuthentication(svc), handler.GetUserInfoHandler(svc))    // lấy thông tin user // done
	router.PUT("api/user/info", middleware.UserAuthentication(svc), handler.UpdateUserInfoHandler(svc)) // cập nhật thông tin user // done
	router.POST("api/refresh-token", handler.RefreshTokenHandler(svc))                                  // refresh token // done

	//topic api
	router.POST("api/topic", handler.CreateTopicHandler(svc))        // đăng kí nhóm + đăng kí đề tài
	router.POST("api/topic/accept", handler.AcceptTopicHandler(svc)) // duyệt đề tài
	router.PUT("api/topic/:id", handler.UpdateTopicHandler(svc))     // sửa đề tài
	router.GET("api/topic", handler.GetTopicHandler(svc))            // lấy danh sách đề tài
	router.GET("api/topic/id", handler.GetTopicByIdHandler(svc))     // lấy đề tài theo id

	//topic registration api
	router.GET("api/topic-registation", handler.GetListTopicRegistationHandler(svc))                                     // lấy danh sách đề tài đề xuất // done
	router.GET("api/topic-registation/:id", handler.GetTopicRegistaionByIDHandler(svc))                                  // lây đề tài đề xuất theo id // done
	router.POST("api/topic-registation/accept", handler.AcceptTopicRegistationHandler(svc))                              // duyệt đề tài đề xuât // done
	router.PUT("api/topic-registation/:id", handler.UpdateTopicRegistationHandler(svc))                                  // cập nhật đề tài đề xuất // done
	router.POST("api/topic-registation", middleware.UserAuthentication(svc), handler.CreateTopicRegistationHandler(svc)) // tạo đề tài đề xuất // done

	//conference
	// lấy danh sách nhóm
	// lấy danh sach thanh vien theo id nhom
	// tao nhom
	// cap nhat thanh vien nhom

	//department api
	router.GET("api/department", handler.GetListDepartmentHandler(svc))
	router.GET("api/department/:id", handler.GetDepartmentByIDHandler(svc))
	router.POST("api/department", handler.CreateDepartmentHandler(svc))

	//faculity api
	router.GET("api/faculity", handler.GetListFaculityHandler(svc))
	router.GET("api/faculity/:id", handler.GetFaculityByIDHandler(svc))
	router.POST("api/faculty", handler.CreateFaculityHandler(svc))
}
