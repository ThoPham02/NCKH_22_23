// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/departments",
				Handler: GetDepartmentsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/department",
				Handler: CreateDepartmentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/faculties",
				Handler: GetFacultiesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/faculty",
				Handler: CreateFacultyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/topics",
				Handler: GetTopicsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/topic/:id",
				Handler: GetTopicHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/topic",
				Handler: CreateTopicHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/topic/:id",
				Handler: UpdateTopicHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/topic-status/:id",
				Handler: UpdateTopicStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/topic-subcommittee",
				Handler: UpdateTopicSubcommitteeHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/topic-student-group/:id",
				Handler: UpdateTopicStudentGroupHandler(serverCtx),
			},
		},
	)
}
