package handler

import (
	"context"
	"github/ThoPham02/research_management/api/logic"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"github/ThoPham02/research_management/core/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListTopicRegistationHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace-id", logger.GenerateTraceID("get-list-topic-registation"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := types.GetListTopicRegistationRequest{}
		err := http_request.BindQueryString(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, nil)
			return
		}

		res, err := logic.GetListTopicRegistationLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}
