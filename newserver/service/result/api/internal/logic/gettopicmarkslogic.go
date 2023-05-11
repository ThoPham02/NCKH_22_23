package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicMarksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicMarksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicMarksLogic {
	return &GetTopicMarksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicMarksLogic) GetTopicMarks(req *types.GetTopicMarksReq) (resp *types.GetTopicMarksRes, err error) {
	// todo: add your logic here and delete this line

	return
}
