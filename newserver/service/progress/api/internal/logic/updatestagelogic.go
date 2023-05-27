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

type UpdateStageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStageLogic {
	return &UpdateStageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStageLogic) UpdateStage(req *types.UpdateStageReq) (resp *types.UpdateStageRes, err error) {
	l.Logger.Info("UpdateStage", req)

	stage, err := l.svcCtx.StageModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateStageRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if stage == nil {
		return &types.UpdateStageRes{
			Result: types.Result{
				Code:    common.STAGE_NOT_EXISTS_ERR_CODE,
				Message: common.STAGE_NOT_EXISTS_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.StageModel.Update(l.ctx, &model.StageTbl{
		Id:          req.ID,
		Name:        stage.Name,
		Description: sql.NullString{Valid: true, String: req.Description},
		EventId:     stage.EventId,
		TimeStart:   sql.NullInt64{Valid: true, Int64: req.TimeStart},
		TimeEnd:     sql.NullInt64{Valid: true, Int64: req.TimeEnd},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateStageRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.UpdateStageRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
