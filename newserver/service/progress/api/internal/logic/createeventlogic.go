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
	var stages []types.Stage
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
	stageID := sync.RandomID()
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID,
		Name:    "Đề Xuất",
		EventId: eventID,
	})
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID + 1,
		Name:    "Đăng ký",
		EventId: eventID,
	})
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID + 2,
		Name:    "Thực hiện",
		EventId: eventID,
	})
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID + 3,
		Name:    "Báo cáo tiến độ",
		EventId: eventID,
	})
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID + 4,
		Name:    "Bộ môn nghiệm thu",
		EventId: eventID,
	})
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID + 5,
		Name:    "Báo cáo tiểu ban",
		EventId: eventID,
	})
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID + 6,
		Name:    "Báo cáo cấp trường",
		EventId: eventID,
	})
	l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:      stageID + 7,
		Name:    "Kết thúc",
		EventId: eventID,
	})

	stagesModel, err := l.svcCtx.StageModel.FindStages(l.ctx, eventID)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateEventRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, stageModel := range stagesModel {
		stages = append(stages, types.Stage{
			ID:      stageModel.Id,
			Name:    stageModel.Name,
			EventID: eventID,
		})
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
			Stages:     stages,
		},
	}, nil
}
