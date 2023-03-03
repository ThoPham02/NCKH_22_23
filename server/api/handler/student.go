package handler

import (
	"database/sql"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetStudentParam struct {
	ID int64 `uri:"id"`
}

func GetStudentHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetStudentParam
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrRequest(err))
			return
		}

		student, err := svc.Store.GetStudent(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, utils.ErrResourceNotFound)
				return
			}
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, student)
	}
}
