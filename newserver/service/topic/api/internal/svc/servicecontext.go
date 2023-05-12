package svc

import (
	"github.com/ThoPham02/research_management/service/topic/api/internal/config"
	"github.com/ThoPham02/research_management/service/topic/model"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	DepartmentModel   model.DepartmentTblModel
	FacultyModel      model.FacultyTblModel
	StudentGroupModel model.StudentGroupTblModel
	TopicModel        model.TopicTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.Name, c.Database.Source)

	return &ServiceContext{
		Config:            c,
		DepartmentModel:   model.NewDepartmentTblModel(conn),
		FacultyModel:      model.NewFacultyTblModel(conn),
		TopicModel:        model.NewTopicTblModel(conn),
		StudentGroupModel: model.NewStudentGroupTblModel(conn),
	}
}
