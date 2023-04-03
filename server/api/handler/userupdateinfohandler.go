package handler

import (
	"context"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/logic"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"github/ThoPham02/research_management/core/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserInfoHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace-id", logger.GenerateTraceID("update-user-info-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.UserInfoResponse{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorf("Failed while parse user info response: %v", err)
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		payload := c.Value(constant.PayloadKey).(*token.Payload)
		err = logic.UpdateUserInfo(payload.UserID, req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, nil)
	}
}
