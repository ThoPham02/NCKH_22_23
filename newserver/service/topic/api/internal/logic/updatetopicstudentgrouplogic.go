package logic

import (
	"context"

	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicStudentGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicStudentGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicStudentGroupLogic {
	return &UpdateTopicStudentGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicStudentGroupLogic) UpdateTopicStudentGroup(req *types.UpdateTopicStudentGroupReq) (resp *types.UpdateTopicStudentGroupRes, err error) {
	// todo: add your logic here and delete this line

	return
}
