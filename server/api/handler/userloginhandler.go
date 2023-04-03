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

func UserLoginHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace-id", logger.GenerateTraceID("user-login-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.UserLoginRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorw("msg", "Failed while parsing user login request", "error", err.Error())
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		res, err := logic.Login(req)
		if err != nil {
			if err.Error() == constant.WrongPasswordErrMsg || err.Error() == constant.UserIsNotExistErrMsg {
				http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
				return
			}
			http_response.ResponseJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}
