package api

import (
	"flag"

	"github.com/ThoPham02/research_management/service/topic/api/internal/config"
	"github.com/ThoPham02/research_management/service/topic/api/internal/handler"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("topic-conf", "etc/topic-api.yaml", "the config file")

type TopicService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewTopicService(server *rest.Server) *TopicService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &TopicService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *TopicService) Start() error {
	return nil
}
