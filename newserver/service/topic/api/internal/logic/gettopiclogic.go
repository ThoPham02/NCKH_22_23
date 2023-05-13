package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicLogic {
	return &GetTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicLogic) GetTopic(req *types.GetTopicReq) (resp *types.GetTopicRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetTopic", req)

	var topicModel *model.TopicTbl
	var topic types.Topic
	topicModel, err = l.svcCtx.TopicModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.GetTopicRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	topic = types.Topic{
		ID:              topicModel.Id,
		Name:            topicModel.Name,
		LectureID:       topicModel.LectureId,
		DepartmentID:    topicModel.DepartmentId,
		Status:          topicModel.Status,
		SubcommitteeID:  topicModel.SubcommitteeId.Int64,
		GroupStudentsID: topicModel.GroupStudentsId.Int64,
		TimeStart:       topicModel.TimeStart.Int64,
		TimeEnd:         topicModel.TimeEnd.Int64,
	}

	return &types.GetTopicRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Topic: topic,
	}, nil
}
