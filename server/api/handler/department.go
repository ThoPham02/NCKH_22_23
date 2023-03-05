package handler

import (
	"database/sql"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetListDepartmentsRequest struct {
	FacultyID int64 `form:"faculty_id"`
}

func GetListDepartmentsHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetListDepartmentsRequest
		if err := ctx.BindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrResponse(err))
			return
		}

		departments, err := svc.Store.ListDepartments(ctx, req.FacultyID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, departments)
	}
}

func GetDepartmentInfoHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := ctx.MustGet("payload_key").(*token.Payload)

		user, err := svc.Store.GetUserByName(ctx, payload.UserName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
		}

		department, err := svc.Store.GetDepartment(ctx, user.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, utils.ErrResourceNotFound)
				return
			}
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, department)
	}
}
