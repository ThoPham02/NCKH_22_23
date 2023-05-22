package svc

import (
	userModel "github.com/ThoPham02/research_management/service/account/model"
	progressModel "github.com/ThoPham02/research_management/service/progress/model"
	subcommitteeModel "github.com/ThoPham02/research_management/service/result/model"
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
	EventModel        progressModel.EventTblModel
	UserModel         userModel.UserTblModel
	SubcommitteeModel subcommitteeModel.SubcommitteeTblModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.Name, c.Database.Source)

	return &ServiceContext{
		Config:            c,
		DepartmentModel:   model.NewDepartmentTblModel(conn),
		FacultyModel:      model.NewFacultyTblModel(conn),
		TopicModel:        model.NewTopicTblModel(conn),
		StudentGroupModel: model.NewStudentGroupTblModel(conn),
		EventModel:        progressModel.NewEventTblModel(conn),
		UserModel:         userModel.NewUserTblModel(conn),
		SubcommitteeModel: subcommitteeModel.NewSubcommitteeTblModel(conn),
	}
}
