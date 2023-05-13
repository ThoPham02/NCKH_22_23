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

type GetTopicReportsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicReportsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicReportsLogic {
	return &GetTopicReportsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicReportsLogic) GetTopicReports(req *types.GetTopicReportsReq) (resp *types.GetTopicReportsRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetTopicReports", req)

	var topicReportsModel []model.TopicReportTbl
	var topicReportModel model.TopicReportTbl
	var topicReports []types.TopicReport
	var topicReport types.TopicReport
	var count int64

	topicReportsModel, err = l.svcCtx.TopicReportModel.FindTopicReports(l.ctx, model.TopicReportConditions{
		StageID: req.StageID,
		EventID: req.EventID,
		TopicID: req.TopicID,
		Limit:   req.Limit,
		Offset:  req.Offset,
	})
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.GetTopicReportsRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil

		}
		l.Logger.Error(err)
		return &types.GetTopicReportsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	count, err = l.svcCtx.TopicReportModel.CountTopicReports(l.ctx, model.TopicReportConditions{
		StageID: req.StageID,
		EventID: req.EventID,
		TopicID: req.TopicID,
		Limit:   req.Limit,
		Offset:  req.Offset,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicReportsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, topicReportModel = range topicReportsModel {
		topicReport = types.TopicReport{
			ID:          topicReportModel.Id,
			TopicID:     topicReportModel.TopicId,
			Description: topicReportModel.Description.String,
			ReportUrl:   topicReportModel.ReportUrl,
			StageID:     topicReportModel.StageId,
		}
		topicReports = append(topicReports, topicReport)
	}

	return &types.GetTopicReportsRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:        count,
		TopicReports: topicReports,
	}, nil
}
