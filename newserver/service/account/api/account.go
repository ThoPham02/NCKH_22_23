package api

import (
	"flag"

	"github.com/ThoPham02/research_management/service/account/api/internal/config"
	"github.com/ThoPham02/research_management/service/account/api/internal/handler"
	"github.com/ThoPham02/research_management/service/account/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("account-conf", "etc/account-api.yaml", "the config file")

type AccountService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewAccountService(server *rest.Server) *AccountService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &AccountService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *AccountService) Start() error {
	return nil
}
