package service

import (
	"github/ThoPham02/research_management/api/db/store"
	"github/ThoPham02/research_management/config"
)

type ServiceContext struct {
	Config config.Config
	Store  store.Store
}

func NewServiceContext(c config.Config, store store.Store) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Store:  store,
	}
}
