package logic

import (
	"context"
	"github/ThoPham02/research_management/api/service"

	"github.com/go-kratos/kratos/v2/log"
)

type Logic struct {
	ctx       context.Context
	svcCtx    *service.ServiceContext
	logHelper *log.Helper
}

func NewLogic(ctx context.Context, svcCtx *service.ServiceContext, logHelper *log.Helper) Logic {
	return Logic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}
