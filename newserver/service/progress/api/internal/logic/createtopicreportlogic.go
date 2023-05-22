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

type CreateTopicReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTopicReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTopicReportLogic {
	return &CreateTopicReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTopicReportLogic) CreateTopicReport(req *types.CreateTopicReportReq) (resp *types.CreateTopicReportRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("CreateTopicReport", req)

	_, err = l.svcCtx.TopicReportModel.Insert(l.ctx, &model.TopicReportTbl{
		Id:      sync.RandomID(),
		TopicId: req.TopicID,
		StageId: req.StageID,
		Description: sql.NullString{
			Valid:  true,
			String: req.Description,
		},
		ReportUrl: req.ReportUrl,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateTopicReportRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.CreateTopicReportRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
