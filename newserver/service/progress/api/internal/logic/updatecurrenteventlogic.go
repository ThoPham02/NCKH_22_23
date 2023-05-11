package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCurrentEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCurrentEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCurrentEventLogic {
	return &UpdateCurrentEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCurrentEventLogic) UpdateCurrentEvent(req *types.UpdateCurrentEventReq) (resp *types.UpdateCurrentEventRes, err error) {
	// todo: add your logic here and delete this line

	return
}
