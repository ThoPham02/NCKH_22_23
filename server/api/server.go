package api

import (
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/config"
	"github/ThoPham02/research_management/db/store"
	"github/ThoPham02/research_management/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	svc        *service.ServiceContext
	TokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(store store.Store) *Server {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config::", err)
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSemmetricKey)
	if err != nil {
		log.Fatal("Can't create a new token::", err)
	}
	router := gin.Default()
	svc := service.NewServiceContext(*config, store)
	RegisterRouter(router, svc)

	server := &Server{
		svc:        svc,
		TokenMaker: tokenMaker,
		Router:     router,
	}

	return server
}
