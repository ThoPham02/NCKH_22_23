package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicReportsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicReportsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicReportsLogic {
	return &GetTopicReportsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicReportsLogic) GetTopicReports(req *types.GetTopicReportsReq) (resp *types.GetTopicReportsRes, err error) {
	// todo: add your logic here and delete this line

	return
}
