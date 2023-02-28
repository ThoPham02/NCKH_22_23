package api

import (
	"github/ThoPham02/research_management/api/handler"
	"github/ThoPham02/research_management/db/store"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(store store.Store, router *gin.Engine) *gin.Engine {
	router.GET("/user/:id", handler.GetUserHandler(store))

	return router
}
