package handler

import (
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateConferenceHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logic := InitLogic(svcCtx, ctx, "create-conference")

		var req = types.CreateConferenceRequest{}
		err := http_request.BindBodyJson(ctx, &req)
		if err != nil {
			http_response.ResponseJSON(ctx, http.StatusBadRequest, err)
			return
		}

		resp, err := logic.CreateConferenceLogic(&req)
		if err != nil {
			http_response.ResponseJSON(ctx, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(ctx, http.StatusOK, resp)
	}
}

func ListConferenceHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logic := InitLogic(svcCtx, ctx, "create-conference")

		resp, err := logic.ListConferenceLogic()
		if err != nil {
			http_response.ResponseJSON(ctx, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(ctx, http.StatusOK, resp)
	}
}
