package handler

import (
	"database/sql"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLecturerInfoHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := ctx.MustGet("payload_key").(*token.Payload)

		user, err := svc.Store.GetUserByName(ctx, payload.UserName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
		}

		lecturer, err := svc.Store.GetLecturer(ctx, user.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, utils.ErrResourceNotFound)
				return
			}
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, lecturer)
	}
}
