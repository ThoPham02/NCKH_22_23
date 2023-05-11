package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFacultyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFacultyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFacultyLogic {
	return &CreateFacultyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFacultyLogic) CreateFaculty(req *types.CreateFacultyReq) (resp *types.CreateFacultyRes, err error) {
	// todo: add your logic here and delete this line

	return
}
