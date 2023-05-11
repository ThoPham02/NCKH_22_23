package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicReportLogic {
	return &GetTopicReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicReportLogic) GetTopicReport(req *types.GetTopicReportsByStageReq) (resp *types.GetTopicReportsByStageRes, err error) {
	// todo: add your logic here and delete this line

	return
}
