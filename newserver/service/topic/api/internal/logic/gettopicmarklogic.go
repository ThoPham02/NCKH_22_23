package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	resultModel "github.com/ThoPham02/research_management/service/result/model"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicMarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicMarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicMarkLogic {
	return &GetTopicMarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicMarkLogic) GetTopicMark(req *types.GetTopicsReq) (resp *types.GetTopicMarkRes, err error) {
	// todo: add your logic here and delete this line
	var data = []types.TopicMark{}
	var mapMark = map[int64]float64{}
	var conditions = model.TopicConditions{
		Search:         req.Search,
		DepartmentID:   req.DepartmentID,
		FacultyID:      req.FacultyID,
		Status:         req.Status,
		LectureID:      req.LectureID,
		StudentID:      req.StudentID,
		IsCurrent:      req.IsCurrent,
		EventID:        req.EventID,
		SubcommitteeID: req.SubcommitteeID,
		TimeStart:      req.TimeStart,
		TimeEnd:        req.TimeEnd,
		Limit:          req.Limit,
		Offset:         req.Offset,
	}
	total, err := l.svcCtx.TopicModel.CountTopics(l.ctx, conditions)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicMarkRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	topicsModel, err := l.svcCtx.TopicModel.FindTopics(l.ctx, conditions)
	l.Logger.Info(err)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.GetTopicMarkRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetTopicMarkRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	topicMark, err := l.svcCtx.TopicMarkModel.FindTopicMarks(l.ctx, resultModel.TopicMarkConditions{
		TopicIDs: []int64{},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicMarkRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, tmp := range topicMark {
		mapMark[tmp.TopicId] += tmp.Point
	}

	subs, err := l.svcCtx.SubcommitteeModel.FindSubcommittee(l.ctx, resultModel.SubcommitteeConditions{
		EventID: 0,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicMarkRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	var mapSub = map[int64]types.Subcommittee{}
	for _, tmp := range subs {
		mapSub[tmp.Id] = types.Subcommittee{
			ID:        tmp.Id,
			Name:      tmp.Name,
			FacultyID: tmp.FacultId.Int64,
			EventID:   tmp.EventId,
			Level:     tmp.Level,
		}
	}

	for _, tmp := range topicsModel {
		data = append(data, types.TopicMark{
			ID:             tmp.Id,
			Name:           tmp.Name,
			SubcommitteeID: mapSub[tmp.Id],
			Mark:           mapMark[tmp.Id],
		})
	}

	return &types.GetTopicMarkRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:     total,
		TopicMark: data,
	}, nil
}
