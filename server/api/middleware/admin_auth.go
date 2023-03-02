package middleware

import (
	"errors"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "Bearer"
	payloadKey              = "payload_key"
)

func AdminAuthentication(svc *service.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResposive(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 || fields[0] != authorizationTypeBearer {
			err := errors.New("invalid authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResposive(err))
			return
		}

		tokenMaker, err := token.NewPasetoMaker(svc.Config.TokenSemmetricKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrResposive(err))
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrResposive(err))
			return
		}

		ctx.Set(payloadKey, payload)

		ctx.Next()
	}
}
