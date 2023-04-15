package handler

import (
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"
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
		// ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("get-list-faclity"))
		// logHelper := logger.NewContextLog(ctx)
		// logic := logic.NewLogic(ctx, svcCtx, logHelper)

		// result, err := logic.GetListFaculityLogic()
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusInternalServerError, err)
		// 	return
		// }
		// http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
