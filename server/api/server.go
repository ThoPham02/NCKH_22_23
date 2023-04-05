package api

import (
	"github/ThoPham02/research_management/api/db/store"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	svc    *service.ServiceContext
	Router *gin.Engine
}

func NewServer(store store.Store) *Server {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config::", err)
	}

	router := gin.Default()
	svc := service.NewServiceContext(*config, store)
	RegisterRouter(router, svc)

	server := &Server{
		svc:    svc,
		Router: router,
	}

	return server
}
