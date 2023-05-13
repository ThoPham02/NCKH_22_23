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

type UpdateTopicStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicStatusLogic {
	return &UpdateTopicStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicStatusLogic) UpdateTopicStatus(req *types.UpdateTopicStatusReq) (resp *types.UpdateTopicStatusRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("UpdateTopicStatus", req)

	var topicModel *model.TopicTbl

	topicModel, err = l.svcCtx.TopicModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.UpdateTopicStatusRes{
				Result: types.Result{
					Code:    common.TOPIC_NOT_EXIST_CODE,
					Message: common.TOPIC_NOT_EXIST_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.UpdateTopicStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.TopicModel.Update(l.ctx, &model.TopicTbl{
		Id:              topicModel.Id,
		Name:            topicModel.Name,
		LectureId:       topicModel.LectureId,
		DepartmentId:    topicModel.DepartmentId,
		Status:          req.Status,
		EventId:         topicModel.EventId,
		SubcommitteeId:  topicModel.SubcommitteeId,
		TimeStart:       topicModel.TimeStart,
		TimeEnd:         topicModel.TimeEnd,
		CashSupport:     topicModel.CashSupport,
		GroupStudentsId: topicModel.GroupStudentsId,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicStatusRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.UpdateTopicStatusRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
