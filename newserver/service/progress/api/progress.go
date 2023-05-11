package api

import (
	"flag"

	"github.com/ThoPham02/research_management/service/progress/api/internal/config"
	"github.com/ThoPham02/research_management/service/progress/api/internal/handler"
	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("progress-conf", "etc/progress-api.yaml", "the config file")

type ProgressService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewProgressService(server *rest.Server) *ProgressService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &ProgressService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *ProgressService) Start() error {
	return nil
}
