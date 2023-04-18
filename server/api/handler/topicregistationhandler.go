package handler

import (
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListTopicRegistationHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "get-list-topic-registation")

		req := types.GetTopicRegistrationsRequest{}
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

func GetTopicRegistaionByIDHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "get-topic-regist-by-id")

		req := types.GetTopicRegistrationByIdRequest{}
		err := http_request.BindUri(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		id, err := GetUriID(c)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		res, err := logic.GetTopicRegistationByIdLogic(id, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func CreateTopicRegistationHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "create-topic-regis")

		req := types.CreateTopicRegistrationRequest{}
		err := http_request.BindBodyJson(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		payload := token.GetPayload(c)
		res, err := logic.CreateTopicRegistationLogic(payload.UserID, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func UpdateTopicRegistationHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "update-topic-regis")

		var req types.UpdateTopicRegistrationRequest
		err := http_request.BindBodyJson(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		res, err := logic.UpdateTopicRegistationLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}
