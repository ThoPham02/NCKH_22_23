package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTopicLogic {
	return &CreateTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTopicLogic) CreateTopic(req *types.CreateTopicReq) (resp *types.CreateTopicRes, err error) {
	// todo: add your logic here and delete this line

	return
}
