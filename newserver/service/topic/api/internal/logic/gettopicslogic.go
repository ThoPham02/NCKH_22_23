package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicsLogic {
	return &GetTopicsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicsLogic) GetTopics(req *types.GetTopicsReq) (resp *types.GetTopicsRes, err error) {
	// todo: add your logic here and delete this line

	return
}
