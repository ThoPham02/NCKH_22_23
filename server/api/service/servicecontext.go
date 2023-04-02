package service

import (
	"github/ThoPham02/research_management/api/db/store"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/config"
)

type ServiceContext struct {
	Config config.Config
	Store  store.Store
	Maker  token.Maker
}

func NewServiceContext(c config.Config, store store.Store) *ServiceContext {
	maker, err := token.NewPasetoMaker(c.TokenSemmetricKey)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Store:  store,
		Maker:  maker,
	}
}
