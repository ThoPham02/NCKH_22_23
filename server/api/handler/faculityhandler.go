package handler

import (
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFaculityByIDHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svc, c, "get-faculity-by-id")

		id, err := GetUriID(c)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		res, err := logic.GetFaculityByIDLogic(id, &types.GetFacultyByIDRequest{})
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func GetListFaculityHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "get-list-faculity")

		req := types.GetFacultysRequest{}
		err := http_request.BindFormData(c, &req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		result, err := logic.GetListFaculityLogic(&req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
