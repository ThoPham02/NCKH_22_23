package main

import (
	"flag"
	"fmt"

	"github.com/ThoPham02/research_management/config"
	accountApi "github.com/ThoPham02/research_management/service/account/api"
	progressApi "github.com/ThoPham02/research_management/service/progress/api"
	resultApi "github.com/ThoPham02/research_management/service/result/api"
	topicApi "github.com/ThoPham02/research_management/service/topic/api"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("server_config", "etc/server.yaml", "the config file")

// @BasePath  /api
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))

	logx.DisableStat()
	defer server.Stop()

	accountService := accountApi.NewAccountService(server)
	accountService.Start()

	topicService := topicApi.NewTopicService(server)
	topicService.Start()

	progressService := progressApi.NewProgressService(server)
	progressService.Start()

	resultService := resultApi.NewResultService(server)
	resultService.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
