package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStagesLogic {
	return &GetStagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStagesLogic) GetStages(req *types.GetStagesReq) (resp *types.GetStagesRes, err error) {
	// todo: add your logic here and delete this line

	return
}
