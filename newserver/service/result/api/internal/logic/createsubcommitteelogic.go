package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubcommitteeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSubcommitteeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubcommitteeLogic {
	return &CreateSubcommitteeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubcommitteeLogic) CreateSubcommittee(req *types.CreateSubcommitteeReq) (resp *types.CreateSubcommitteeRes, err error) {
	// todo: add your logic here and delete this line

	return
}
