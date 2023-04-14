package handler

import (
	"context"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/logic"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"github/ThoPham02/research_management/core/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListDepartmentHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("get-list-department-by-faculity"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		var req types.GetListDepartmentByFaculityRequest
		err := http_request.BindQueryString(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		result, err := logic.GetListDepartmentLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}

func GetDepartmentByIDHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("get-list-department-by-id"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		var req types.GetDepartmentByIDRequest
		err := http_request.BindUri(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		result, err := logic.GetDepartmentByIDLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
