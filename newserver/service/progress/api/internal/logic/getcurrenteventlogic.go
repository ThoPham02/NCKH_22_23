package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentEventLogic {
	return &GetCurrentEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentEventLogic) GetCurrentEvent(req *types.GetCurrentEventReq) (resp *types.GetCurrentEventRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetCurrentEvent, req")

	event, err := l.svcCtx.EventModel.FindCurrentEvent(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetCurrentEventRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.GetCurrentEventRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Events: types.Event{
			ID:         event.Id,
			Name:       event.Name,
			SchoolYear: event.SchoolYear.String,
			IsCurrent:  event.IsCurrent.Int64,
		},
	}, nil
}
