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

// GetUserInfoHandler godoc
//
//	@Summary	get_user_info
//	@Schemes
//	@Description	get_user_info
//	@Tags			user
//	@Success		200
//	@Failure		400
//	@securityDefinitions.basic	BasicAuth
//	@Router			/api/user/info [GET]
func GetUserInfoHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace-id", logger.GenerateTraceID("get-user-info-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		payload := c.Value(constant.PayloadKey).(*token.Payload)
		info, err := logic.GetUserInfo(payload.UserID)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, info)
	}
}
