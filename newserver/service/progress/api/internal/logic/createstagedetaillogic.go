package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStageDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStageDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStageDetailLogic {
	return &CreateStageDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStageDetailLogic) CreateStageDetail(req *types.CreateStageDetailReq) (resp *types.CreateStageDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
