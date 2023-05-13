package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTopicReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTopicReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTopicReportLogic {
	return &CreateTopicReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTopicReportLogic) CreateTopicReport(req *types.CreateTopicReportReq) (resp *types.CreateTopicReportRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("CreateTopicReport", req)

	return
}
