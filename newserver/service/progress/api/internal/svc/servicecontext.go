package svc

import (
	"github.com/ThoPham02/research_management/service/progress/api/internal/config"
	"github.com/ThoPham02/research_management/service/progress/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	EventModel        model.EventTblModel
	LibraryModel      model.LibraryTblModel
	NotificationModel model.NotificationTblModel
	StageModel        model.StageTblModel
	TopicReportModel  model.TopicReportTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.Name, c.Database.Source)
	return &ServiceContext{
		Config:            c,
		EventModel:        model.NewEventTblModel(conn),
		LibraryModel:      model.NewLibraryTblModel(conn),
		NotificationModel: model.NewNotificationTblModel(conn),
		StageModel:        model.NewStageTblModel(conn),
		TopicReportModel:  model.NewTopicReportTblModel(conn),
	}
}
