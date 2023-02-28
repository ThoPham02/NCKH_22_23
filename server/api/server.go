package api

import (
	"github/ThoPham02/research_management/db/store"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Store  store.Store
	Router *gin.Engine
}

func NewServer(store store.Store) *Server {
	server := &Server{Store: store}
	router := RegisterRouter(store, gin.Default())

	server.Router = router

	return server
}
