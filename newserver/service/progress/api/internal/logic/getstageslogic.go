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

type GetStagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStagesLogic {
	return &GetStagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStagesLogic) GetStages(req *types.GetStagesReq) (resp *types.GetStagesRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetStages", req)

	var stagesModel []model.StageTbl
	var stageModel model.StageTbl
	var stages []types.Stage
	var stage types.Stage

	stagesModel, err = l.svcCtx.StageModel.FindStages(l.ctx, req.EventID, req.FacultyID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.GetStagesRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetStagesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, stageModel = range stagesModel {
		stage = types.Stage{
			ID:          stageModel.Id,
			Name:        stageModel.Name,
			Description: stageModel.Description.String,
			FacultyID:   stageModel.FacultyId,
			Url:         stageModel.Url.String,
			EventID:     stageModel.EventId,
			TimeStart:   stageModel.TimeStart,
			TimeEnd:     stageModel.TimeEnd,
		}
		stages = append(stages, stage)
	}

	return &types.GetStagesRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:  int64(len(stages)),
		Stages: stages,
	}, nil
}
