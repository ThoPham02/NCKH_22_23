package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFacultiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFacultiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFacultiesLogic {
	return &GetFacultiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFacultiesLogic) GetFaculties(req *types.GetFacultiesReq) (resp *types.GetFacultiesRes, err error) {
	// todo: add your logic here and delete this line

	return
}
