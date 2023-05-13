package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStageLogic {
	return &CreateStageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStageLogic) CreateStage(req *types.CreateStageReq) (resp *types.CreateStageRes, err error) {
	// todo: add your logic here and delete this line

	return
}
