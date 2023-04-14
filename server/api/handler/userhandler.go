package handler

import (
	"context"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/logic"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/core/http_request"
	"github/ThoPham02/research_management/core/http_response"
	"github/ThoPham02/research_management/core/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("user-login-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.UserLoginRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorw("msg", "Failed while parsing user login request", "error", err.Error())
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		res, err := logic.Login(req)
		if err != nil {
			if err.Error() == constant.WrongPasswordErrMsg || err.Error() == constant.UserIsNotExistErrMsg {
				http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
				return
			}
			http_response.ResponseJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, res)
	}
}

func UserRegisterHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), constant.TraceIDKey, logger.GenerateTraceID("user-register-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.UserRegisterRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorf("Failed while parsing user register request, error: %s", err.Error())
			http_response.ResponseJSON(c, http.StatusBadRequest, nil)
			return
		}

		err = logic.Register(req)
		if err != nil {
			if err.Error() == constant.InputValidationErrMsg {
				http_response.ResponseJSON(c, http.StatusBadRequest, nil)
				return
			}
			http_response.ResponseJSON(c, http.StatusInternalServerError, nil)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, nil)
	}
}

func UpdateUserInfoHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("update-user-info-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.UserInfoResponse{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorf("Failed while parse user info response: %v", err)
			http_response.ResponseJSON(c, http.StatusBadRequest, err)
			return
		}

		payload := c.Value(constant.PayloadKey).(*token.Payload)
		err = logic.UpdateUserInfo(payload.UserID, req)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, nil)
	}
}

func GetUserInfoHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("get-user-info-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		payload := c.Value(constant.PayloadKey).(*token.Payload)
		info, err := logic.GetUserInfo(payload.UserID)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusInternalServerError, err)
			return
		}

		http_response.ResponseJSON(c, http.StatusOK, info)
	}
}

func RefreshTokenHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constant.TraceIDKey, logger.GenerateTraceID("refresh-token-api"))
		logHelper := logger.NewContextLog(ctx)
		logic := logic.NewLogic(ctx, svcCtx, logHelper)

		req := &types.RefreshTokenRequest{}
		err := http_request.BindBodyJson(c, req)
		if err != nil {
			logHelper.Errorf("Failed while get refresh token: %v", err)
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
