package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubcommitteeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubcommitteeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubcommitteeLogic {
	return &GetSubcommitteeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubcommitteeLogic) GetSubcommittee(req *types.GetSubcommitteesReq) (resp *types.GetSubcommitteesRes, err error) {
	// todo: add your logic here and delete this line

	return
}
