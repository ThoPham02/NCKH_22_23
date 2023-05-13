package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"
	"github.com/ThoPham02/research_management/service/progress/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetEventsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEventsLogic {
	return &GetEventsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEventsLogic) GetEvents(req *types.GetEventsReq) (resp *types.GetEventsRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetEvents", req)

	var eventsModel []model.EventTbl
	var eventModel model.EventTbl
	var events []types.Event
	var event types.Event

	eventsModel, err = l.svcCtx.EventModel.FindEvents(l.ctx)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.GetEventsRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetEventsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, eventModel = range eventsModel {
		event = types.Event{
			ID:         eventModel.Id,
			Name:       eventModel.Name,
			SchoolYear: eventModel.SchoolYear.String,
			IsCurrent:  eventModel.IsCurrent.Int64,
		}
		events = append(events, event)
	}

	return &types.GetEventsRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Events: events,
	}, nil
}
