package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"
	"github.com/ThoPham02/research_management/service/result/model"
	topicModel "github.com/ThoPham02/research_management/service/topic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResultLogic {
	return &GetResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetResultLogic) GetResult(req *types.GetResultReq) (resp *types.GetResultRes, err error) {
	// todo: add your logic here and delete this line

	l.Logger.Info("GetResult ", req)

	var topicsModel []topicModel.TopicTbl
	var topics []types.Topic
	var topicMark types.MarkDetail
	var marks []model.TopicMarkTbl
	var topicMarkMap = map[int64][]types.MarkDetail{}
	var total int64

	topicsModel, err = l.svcCtx.TopicModel.FindTopics(l.ctx, topicModel.TopicConditions{
		Search:         req.Search,
		EventID:        req.EventID,
		FacultyID:      req.FacultyID,
		SubcommitteeID: req.SubcommitteeID,
		Limit:          req.Limit,
		Offset:         req.Offset,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetResultRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if len(topicsModel) == 0 {
		return &types.GetResultRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	total, err = l.svcCtx.TopicModel.CountTopics(l.ctx, topicModel.TopicConditions{
		Search:         req.Search,
		EventID:        req.EventID,
		FacultyID:      req.FacultyID,
		SubcommitteeID: req.SubcommitteeID,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetResultRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	marks, err = l.svcCtx.TopicMarkModel.FindTopicMarks(l.ctx, model.TopicMarkConditions{})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetResultRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, mark := range marks {
		topicMark = types.MarkDetail{
			MarkID:    mark.Id,
			Comment:   mark.Comment.String,
			LectureID: mark.LectureId,
			Point:     mark.Point,
		}
		topicMarkMap[mark.TopicId] = append(topicMarkMap[mark.TopicId], topicMark)
	}

	for _, topicModel := range topicsModel {
		topics = append(topics, types.Topic{
			TopicID:   topicModel.Id,
			TopicName: topicModel.Name,
			TopicMark: topicMarkMap[topicModel.Id],
		})
	}

	return &types.GetResultRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Topics: topics,
		Total:  int(total),
	}, nil
}
