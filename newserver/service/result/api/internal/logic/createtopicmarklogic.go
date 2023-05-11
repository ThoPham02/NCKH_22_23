package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTopicMarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTopicMarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTopicMarkLogic {
	return &CreateTopicMarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTopicMarkLogic) CreateTopicMark(req *types.CreateTopicMarkReq) (resp *types.CreateTopicMarkRes, err error) {
	// todo: add your logic here and delete this line

	return
}
