// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/ThoPham02/research_management/service/account/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/find",
				Handler: GetUsersHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/reset",
				Handler: ResetPasswordHandler(serverCtx),
			},
		},
	)
}
