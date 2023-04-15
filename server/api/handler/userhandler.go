package handler

import (
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "user-login")

		req := &types.UserLoginRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		res, err := logic.UserLogin(req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func UserRegisterHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "user-register")

		req := &types.UserRegisterRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, nil)
			return
		}

		resp, err := logic.Register(req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, resp)
	}
}

func UpdateUserInfoHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "update-user")

		req := &types.UpdateUserInfoRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		payload := token.GetPayload(c)
		data, err := logic.UpdateUserInfo(payload.UserID, req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, data)
	}
}

func GetUserInfoHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "get-userinfo")

		payload := c.Value(constant.PayloadKey).(*token.Payload)
		info, err := logic.GetUserInfo(payload.UserID, &types.GetUserInfoRequest{})
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, info)
	}
}

func RefreshTokenHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := InitLogic(svcCtx, c, "refresh-token")

		req := &types.UserRefreshTokenRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		res, err := logic.RefreshToken(req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}
