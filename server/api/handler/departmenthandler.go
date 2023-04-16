package handler

import (
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListDepartmentHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "get-list-department")

		req := types.GetDepartmentsRequest{}
		err := http_request.BindQueryString(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		result, err := logic.GetListDepartmentLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}

func GetDepartmentByIDHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "get-department")

		id, err := GetUriID(c)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		result, err := logic.GetDepartmentByIDLogic(id, &types.GetDepartmentByIDRequest{})
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}

func CreateDepartmentHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "create-faculity")

		req := types.CreateDepartmentRequest{}
		err := http_request.BindBodyJson(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		result, err := logic.CreateDepartmentLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
