package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicStatusLogic {
	return &UpdateTopicStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicStatusLogic) UpdateTopicStatus(req *types.UpdateTopicStatusReq) (resp *types.UpdateTopicStatusRes, err error) {
	// todo: add your logic here and delete this line

	return
}
