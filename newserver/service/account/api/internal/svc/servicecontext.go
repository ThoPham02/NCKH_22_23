package svc

import (
	"github.com/ThoPham02/research_management/service/account/api/internal/config"
	"github.com/ThoPham02/research_management/service/account/model"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserTblModel
	Corsmiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserTblModel(sqlx.NewSqlConn(c.Database.Name, c.Database.Source)),
	}
}
