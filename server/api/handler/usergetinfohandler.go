package handler

import (
	"context"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/logic"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/core/http_response"
	"github/ThoPham02/research_management/core/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfoHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace-id", logger.GenerateTraceID("get-user-info-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		test := c.Value(constant.PayloadKey).(*token.Payload)
		info, err := logic.GetUserInfo(test.UserID)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, info)
	}
}
