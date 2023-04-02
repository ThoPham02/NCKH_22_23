package handler

import (
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetListTopicsRequest struct {
	Page     int32 `form:"page"`
	PageSize int32 `form:"page_size"`
}

func GetListTopicsHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetListTopicsRequest

		if err := ctx.BindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrResponse(err))
			return
		}

		topic, err := svc.Store.ListTopics(ctx, db.ListTopicsParams{
			Limit:  req.PageSize,
			Offset: req.PageSize * (req.Page - 1),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, topic)
	}
}
