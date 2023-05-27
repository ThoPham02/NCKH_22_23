package logic

import (
	"context"
	"database/sql"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"
	"github.com/ThoPham02/research_management/service/progress/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelEventLogic {
	return &CancelEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelEventLogic) CancelEvent(req *types.CancelEventReq) (resp *types.CancelEventRes, err error) {
	// todo: add your logic here and delete this line

	l.Logger.Info("CancelEvent", req)

	event, err := l.svcCtx.EventModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.CancelEventRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.EventModel.Update(l.ctx, &model.EventTbl{
		Id:         event.Id,
		Name:       event.Name,
		SchoolYear: event.SchoolYear,
		IsCurrent:  sql.NullInt64{Valid: true, Int64: 0},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CancelEventRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.CancelEventRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
