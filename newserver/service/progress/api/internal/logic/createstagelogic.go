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

type CreateStageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStageLogic {
	return &CreateStageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStageLogic) CreateStage(req *types.CreateStageReq) (resp *types.CreateStageRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("CreateStage", req)

	_, err = l.svcCtx.StageModel.Insert(l.ctx, &model.StageTbl{
		Id:   sync.RandomID(),
		Name: req.Name,
		Description: sql.NullString{
			Valid:  true,
			String: req.Description,
		},
		Url: sql.NullString{
			Valid:  true,
			String: req.Url,
		},
		EventId:   req.EventID,
		TimeStart: req.TimeStart,
		TimeEnd:   req.TimeEnd,
		FacultyId: req.FacultyID,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateStageRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.CreateStageRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
