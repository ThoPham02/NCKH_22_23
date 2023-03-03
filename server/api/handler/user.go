package handler

import (
	"database/sql"
	"fmt"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := ctx.MustGet("payload_key").(*token.Payload)
		user, err := svc.Store.GetUserByName(ctx, payload.UserName)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, utils.ErrResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type NewUserResponse struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Permission int64  `json:"permission"`
}

type LoginUserResponse struct {
	AccessToken string          `json:"access_token"`
	User        NewUserResponse `json:"user"`
}

func UserRegister(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginUserRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrRequest(err))
			return
		}
		user, err := svc.Store.GetUserByName(ctx, req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, utils.ErrResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}

		isUser := utils.ComparePassword(req.Password, user.Password)
		if !isUser {
			ctx.JSON(http.StatusBadRequest, utils.ErrResponse(fmt.Errorf("%s is wrong password", req.Password)))
			return
		}

		userResponse := NewUserResponse{
			Username:   user.Username,
			Email:      user.Email,
			Permission: user.Permission,
		}

		tokenMaker, err := token.NewPasetoMaker(svc.Config.TokenSemmetricKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}

		accessToken, err := tokenMaker.CreateToken(user.Username, svc.Config.AccessTokenDuration)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(err))
			return
		}

		res := LoginUserResponse{
			AccessToken: accessToken,
			User:        userResponse,
		}
		ctx.JSON(http.StatusOK, res)
	}
}
