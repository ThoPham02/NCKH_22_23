package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	userModel "github.com/ThoPham02/research_management/service/account/model"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicsLogic {
	return &GetTopicsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicsLogic) GetTopics(req *types.GetTopicsReq) (resp *types.GetTopicsRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetTopics", req)

	var topicsModel []model.TopicTbl
	var topicModel model.TopicTbl
	var topics []types.Topic
	var topic types.Topic
	var total int64
	var lectureMap = map[int64]types.LectureInfo{}
	var conditions = model.TopicConditions{
		Search:         req.Search,
		DepartmentID:   req.DepartmentID,
		FacultyID:      req.FacultyID,
		Status:         req.Status,
		LectureID:      req.LectureID,
		EventID:        req.EventID,
		SubcommitteeID: req.SubcommitteeID,
		TimeStart:      req.TimeStart,
		TimeEnd:        req.TimeEnd,
		Limit:          req.Limit,
		Offset:         req.Offset,
	}

	topicsModel, err = l.svcCtx.TopicModel.FindTopics(l.ctx, conditions)
	l.Logger.Info(err)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.GetTopicsRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetTopicsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	lectures, err := l.svcCtx.UserModel.FindUserByCondition(l.ctx, userModel.UserCondition{
		Role: 2,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, tmp := range lectures {
		lectureMap[tmp.Id] = types.LectureInfo{
			ID:     tmp.Id,
			Name:   tmp.Name,
			Email:  tmp.Email.String,
			Phone:  tmp.Phone.String,
			Degree: common.MapDegree[tmp.Degree],
		}
	}

	total, err = l.svcCtx.TopicModel.CountTopics(l.ctx, conditions)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, topicModel = range topicsModel {
		topic = types.Topic{
			ID:           topicModel.Id,
			Name:         topicModel.Name,
			LectureInfo:  lectureMap[topicModel.LectureId],
			DepartmentID: topicModel.DepartmentId,
			Status:       topicModel.Status,
		}
		topics = append(topics, topic)
	}

	return &types.GetTopicsRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:  total,
		Topics: topics,
	}, nil
}
