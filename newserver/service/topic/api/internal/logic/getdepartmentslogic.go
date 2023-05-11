package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepartmentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentsLogic {
	return &GetDepartmentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepartmentsLogic) GetDepartments(req *types.GetDepartmentsReq) (resp *types.GetDepartmentsRes, err error) {
	// todo: add your logic here and delete this line

	return
}
