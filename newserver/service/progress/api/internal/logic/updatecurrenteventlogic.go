package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCurrentEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCurrentEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCurrentEventLogic {
	return &UpdateCurrentEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCurrentEventLogic) UpdateCurrentEvent(req *types.UpdateCurrentEventReq) (resp *types.UpdateCurrentEventRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("UpdateCurrentEvent", req)

	err = l.svcCtx.EventModel.UpdateCurrentEvent(l.ctx, req.EventID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateCurrentEventRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.UpdateCurrentEventRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
