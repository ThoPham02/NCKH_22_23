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

func UserRegisterHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), "trace-id", logger.GenerateTraceID("user-register-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.UserRegisterRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorf("Failed while parsing user register request, error: %s", err.Error())
			http_response.ResponseJSON(c, http.StatusBadRequest, nil)
			return
		}

		err = logic.Register(req)
		if err != nil {
			if err.Error() == constant.InputValidationErrMsg {
				http_response.ResponseJSON(c, http.StatusBadRequest, nil)
				return
			}
			http_response.ResponseJSON(c, http.StatusInternalServerError, nil)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, nil)
	}
}
