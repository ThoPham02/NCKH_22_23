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

// RefreshTokenHandler godoc
//
//	@Summary	refresh_token
//	@Schemes
//	@Description	refresh_token
//	@Tags			user
//	@Param			RefreshToken	body		types.RefreshTokenRequest	true	"refresh_token"
//	@Success		200				{object}	types.AccessTokenResponse
//	@Failure		400
//	@Security		ApiAuthKey
//	@Router			/api/refresh-token [GET]
func RefreshTokenHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace-id", logger.GenerateTraceID("refresh-token-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.RefreshTokenRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorf("Failed while get refresh token: %v", err)
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		res, err := logic.RefreshToken(req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, res)

	}
}
