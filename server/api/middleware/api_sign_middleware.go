package middleware

import (
	"github.com/gin-gonic/gin"
)

// const (
// 	TIME_EXPIRED = 30
// )

// type paramSign struct {
// 	stime int64
// 	nonce string
// 	sign  string
// }

// func getParamSign(url string) (*paramSign, error) {
// 	var paramSign *paramSign

// 	params := strings.Split(url, "?")
// 	if len(params) != 2 {
// 		err := errors.New("")
// 		return nil, err
// 	}

// 	return paramSign, nil
// }

func ApiSignMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// url := ctx.Request.URL.String()
		// params, err := getParamSign(url)
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// timeNow := time.Now().UnixNano()
		// if (timeNow-params.stime)/1000 > TIME_EXPIRED {
		// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Time Expired"})
		// 	return
		// }

		// if params.nonce{

		// }

		ctx.Next()
	}
}
