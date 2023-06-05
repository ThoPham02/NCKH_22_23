package svc

import (
	accountModel "github.com/ThoPham02/research_management/service/account/model"
	"github.com/ThoPham02/research_management/service/result/api/internal/config"
	"github.com/ThoPham02/research_management/service/result/model"
	topicModel "github.com/ThoPham02/research_management/service/topic/model"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	SubcommitteeModel model.SubcommitteeTblModel
	GroupModel        model.GroupTblModel
	TopicMarkModel    model.TopicMarkTblModel
	TopicModel        topicModel.TopicTblModel
	UserModel         accountModel.UserTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.Name, c.Database.Source)
	return &ServiceContext{
		Config:            c,
		GroupModel:        model.NewGroupTblModel(conn),
		SubcommitteeModel: model.NewSubcommitteeTblModel(conn),
		TopicMarkModel:    model.NewTopicMarkTblModel(conn),
		TopicModel:        topicModel.NewTopicTblModel(conn),
		UserModel:         accountModel.NewUserTblModel(conn),
	}
}
