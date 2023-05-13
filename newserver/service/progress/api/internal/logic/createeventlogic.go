package logic

import (
	"context"
	"database/sql"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"
	"github.com/ThoPham02/research_management/service/progress/model"
	"github.com/ThoPham02/research_management/sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEventLogic {
	return &CreateEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEventLogic) CreateEvent(req *types.CreateEventReq) (resp *types.CreateEventRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("CreateEvent", req)

	var eventID int64 = sync.RandomID()
	_, err = l.svcCtx.EventModel.Insert(l.ctx, &model.EventTbl{
		Id:   eventID,
		Name: req.Name,
		SchoolYear: sql.NullString{
			Valid:  req.SchoolYear != "",
			String: req.SchoolYear,
		},
		IsCurrent: sql.NullInt64{
			Valid: true,
			Int64: 0,
		},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateEventRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.EventModel.UpdateCurrentEvent(l.ctx, eventID)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateEventRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.CreateEventRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		CurrentEvent: types.Event{
			ID:         eventID,
			Name:       req.Name,
			SchoolYear: req.SchoolYear,
			IsCurrent:  1,
		},
	}, nil
}
