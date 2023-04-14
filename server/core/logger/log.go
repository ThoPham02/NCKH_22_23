package logger

import (
	"context"
	"fmt"
	"github/ThoPham02/research_management/api/constant"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

func Trace() log.Valuer {
	return func(ctx context.Context) interface{} {
		s := ctx.Value(constant.TraceIDKey).(string)
		return s
	}
}

func NewContextLog(ctx context.Context) *log.Helper {
	logger := log.With(log.NewStdLogger(os.Stdout),
		constant.TraceIDKey, Trace(),
	)
	logger = log.With(logger, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)
	return log.NewHelper(logger).WithContext(ctx)
}

func GenerateTraceID(serviceName string) string {
	now := time.Now()
	traceID := fmt.Sprintf("%s-%d", serviceName, now.UnixNano())
	return traceID
}
