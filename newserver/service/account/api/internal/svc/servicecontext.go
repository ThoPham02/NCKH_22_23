package svc

import (
	"github.com/ThoPham02/research_management/service/account/api/internal/config"
	"github.com/ThoPham02/research_management/service/account/model"
	progressModel "github.com/ThoPham02/research_management/service/progress/model"
	resultModel "github.com/ThoPham02/research_management/service/result/model"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	UserModel         model.UserTblModel
	EventModel        progressModel.EventTblModel
	SubcommitteeModel resultModel.SubcommitteeTblModel
	GroupModel        resultModel.GroupTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserModel:         model.NewUserTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.Source)),
		EventModel:        progressModel.NewEventTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.Source)),
		SubcommitteeModel: resultModel.NewSubcommitteeTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.Source)),
		GroupModel:        resultModel.NewGroupTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.Source)),
	}
}
