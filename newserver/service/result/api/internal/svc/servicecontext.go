package svc

import (
	"github.com/ThoPham02/research_management/service/result/api/internal/config"
	"github.com/ThoPham02/research_management/service/result/model"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	SubcommitteeModel model.SubcommitteeTblModel
	GroupModel        model.GroupTblModel
	TopicMark         model.TopicMarkTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.Name, c.Database.Source)
	return &ServiceContext{
		Config:            c,
		GroupModel:        model.NewGroupTblModel(conn),
		SubcommitteeModel: model.NewSubcommitteeTblModel(conn),
		TopicMark:         model.NewTopicMarkTblModel(conn),
	}
}
