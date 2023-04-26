package handler

import (
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopicByIdHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svc, c, "get-topic")

		id, err := GetUriID(c)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		resp, err := logic.GetTopicByIdLogic(id)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, resp)
	}
}

func GetTopicHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svc, c, "get-topic")

		request := types.GetTopicRequest{}
		err := http_request.BindQueryString(c, &request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		resp, err := logic.GetTopicLogic(&request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, resp)
	}
}

func UpdateGroupTopicHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svc, c, "update-topic")

		id, err := GetUriID(c)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		req := types.UpdateTopicRequest{}
		err = http_request.BindBodyJson(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		resp, err := logic.UpdateGroupTopic(id, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, resp)
	}
}

func CreateTopicHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svc, c, "create-topic")
		req := types.CreateTopicRequest{}
		err := http_request.BindBodyJson(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		resp, err := logic.CreateTopicLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, resp)
	}
}

func AcceptTopicHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svc, c, "accept-topic")

		req := types.AcceptTopicRequest{}
		err := http_request.BindBodyJson(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		id, err := GetUriID(c)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		resp, err := logic.AcceptTopicLogic(id, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, resp)
	}
}
