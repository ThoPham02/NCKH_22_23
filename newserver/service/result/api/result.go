package api

import (
	"flag"

	"github.com/ThoPham02/research_management/service/result/api/internal/config"
	"github.com/ThoPham02/research_management/service/result/api/internal/handler"
	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("result-conf", "etc/result-api.yaml", "the config file")

type ResultService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewResultService(server *rest.Server) *ResultService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &ResultService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *ResultService) Start() error {
	return nil
}
