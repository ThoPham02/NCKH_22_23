package middleware

import (
	"errors"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getPayload(ctx *gin.Context, svc *service.ServiceContext) *token.Payload {
	authorizationHeader := ctx.GetHeader(constant.AuthorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(err))
		return nil
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) != 2 || fields[0] != constant.AuthorizationTypeBearer {
		err := errors.New("invalid authorization header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(err))
		return nil
	}

	tokenMaker, err := token.NewPasetoMaker(svc.Config.TokenSemmetricKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrResponse(err))
		return nil
	}
	accessToken := fields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(err))
		return nil
	}

	return payload
}

func AdminAuthentication(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := getPayload(ctx, svc)
		if payload != nil {
			if payload.TypeAccount != constant.AdminType {
				err := errors.New("permission denied")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(err))
				return
			}
			ctx.Set(constant.PayloadKey, payload)
		}
		ctx.Next()
	}
}

func UserAuthentication(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := getPayload(ctx, svc)
		if payload != nil {
			ctx.Set(constant.PayloadKey, payload)
		}
		ctx.Next()
	}
}
