package handler

import (
	"github/ThoPham02/research_management/api/service"

	"github.com/gin-gonic/gin"
)

func GetListTopicRegistationHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("get-list-topic-registation"))
		// logHelper := logger.NewContextLog(ctx)
		// logic := logic.NewLogic(ctx, svcCtx, logHelper)

		// req := types.GetListTopicRegistationRequest{}
		// err := http_request.BindQueryString(c, &req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusBadRequest, nil)
		// 	return
		// }

		// res, err := logic.GetListTopicRegistationLogic(&req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusInternalServerError, err)
		// 	return
		// }
		// http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func GetTopicRegistaionByIDHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("get-topic-registation-by-id"))
		// logHelper := logger.NewContextLog(ctx)
		// logic := logic.NewLogic(ctx, svcCtx, logHelper)

		// req := types.GetTopicRegistationByIDRequest{}
		// err := http_request.BindUri(c, &req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusBadRequest, err)
		// 	return
		// }

		// res, err := logic.GetTopicRegistationByIdLogic(&req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusInternalServerError, err)
		// 	return
		// }
		// http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func CreateTopicRegistationHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("create-topic-registation"))
		// logHelper := logger.NewContextLog(ctx)
		// logic := logic.NewLogic(ctx, svcCtx, logHelper)

		// var req types.CreateTopicRegistationRequest
		// err := http_request.BindBodyJson(c, &req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusBadRequest, err)
		// 	return
		// }

		// res, err := logic.CreateTopicRegistation(&req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusInternalServerError, err)
		// 	return
		// }
		// http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func UpdateTopicRegistationHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("update-topic-registation"))
		// logHelper := logger.NewContextLog(ctx)
		// logic := logic.NewLogic(ctx, svcCtx, logHelper)

		// var req types.UpdateTopicRegistationRequest
		// err := http_request.BindBodyJson(c, &req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusBadRequest, err)
		// 	return
		// }

		// res, err := logic.UpdateTopicRegistationLogic(&req)
		// if err != nil {
		// 	http_response.ResponseJSON(c, http.StatusInternalServerError, err)
		// 	return
		// }
		// http_response.ResponseJSON(c, http.StatusOK, res)
	}
}
