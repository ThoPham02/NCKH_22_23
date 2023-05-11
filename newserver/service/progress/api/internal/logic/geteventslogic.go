package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEventsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEventsLogic {
	return &GetEventsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEventsLogic) GetEvents(req *types.GetEventsReq) (resp *types.GetEventsRes, err error) {
	// todo: add your logic here and delete this line

	return
}
