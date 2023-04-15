package handler

import (
	"context"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/logic"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/logger"

	"github.com/gin-gonic/gin"
)

func InitLogic(svcCtx *service.ServiceContext, c *gin.Context, serviceName string) logic.Logic {
	ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID(serviceName))
	logHelper := logger.NewContextLog(ctx)
	return logic.NewLogic(ctx, svcCtx, logHelper)
}

func GetUriID(c *gin.Context) (int32, error) {
	id := types.IDUri{}
	err := http_request.BindUri(c, &id)
	if err != nil {
		return 0, err
	}
	return id.ID, nil
}
