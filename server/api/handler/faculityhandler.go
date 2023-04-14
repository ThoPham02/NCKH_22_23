package handler

import (
	"github/ThoPham02/research_management/api/service"

	"github.com/gin-gonic/gin"
)

func GetFaculityByIDHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("get-faculity-by-id"))
		// logHelper := logger.NewContextLog(ctx)
		// logic := logic.NewLogic(ctx, svc, logHelper)

		// req := types.GetFaculityByIDRequest{}
		// err := http_request.BindUri(c, &req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusBadRequest, nil)
		// 	return
		// }

		// result, err := logic.GetFaculityByIDLogic(&req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusInternalServerError, err)
		// 	return
		// }
		// http_response.ResponseJSON(c, http.StatusOK, result)
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
