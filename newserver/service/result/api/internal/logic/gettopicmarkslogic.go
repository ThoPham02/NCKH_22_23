package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"
	"github.com/ThoPham02/research_management/service/result/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicMarksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicMarksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicMarksLogic {
	return &GetTopicMarksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicMarksLogic) GetTopicMarks(req *types.GetTopicMarksReq) (resp *types.GetTopicMarksRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetTopicMarks", req)

	var topicMarksModel []model.TopicMarkTbl
	var topicMarkModel model.TopicMarkTbl
	var topicMarks []types.TopicMark
	var topicMark types.TopicMark

	topicMarksModel, err = l.svcCtx.TopicMarkModel.FindTopicMarks(l.ctx, model.TopicMarkConditions{})
	if err != nil {
		if err == model.ErrNotFound {
			return &types.GetTopicMarksRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetTopicMarksRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	for _, topicMarkModel = range topicMarksModel {
		topicMark = types.TopicMark{
			ID:        topicMark.ID,
			TopicID:   topicMarkModel.TopicId,
			LectureID: topicMarkModel.LectureId,
			Point:     topicMarkModel.Point,
			Comment:   topicMarkModel.Comment.String,
			Url:       topicMarkModel.Url.String,
		}
		topicMarks = append(topicMarks, topicMark)
	}

	return &types.GetTopicMarksRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:      int64(len(topicMarks)),
		TopicMarks: topicMarks,
	}, nil
}
