package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	progressModel "github.com/ThoPham02/research_management/service/progress/model"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"
	"github.com/ThoPham02/research_management/sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicStudentGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicStudentGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicStudentGroupLogic {
	return &UpdateTopicStudentGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicStudentGroupLogic) UpdateTopicStudentGroup(req *types.UpdateTopicStudentGroupReq) (resp *types.UpdateTopicStudentGroupRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("UpdateTopicStudentGroup", req)
	var currentEvent *progressModel.EventTbl
	var groupID int64 = sync.RandomID()
	var topic *model.TopicTbl

	currentEvent, err = l.svcCtx.EventModel.FindCurrentEvent(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicStudentGroupRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	topic, err = l.svcCtx.TopicModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicStudentGroupRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	valid, err := l.svcCtx.StudentGroupModel.CheckStudentValid(l.ctx, req.StudentID, currentEvent.Id)
	if !valid || err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicStudentGroupRes{
			Result: types.Result{
				Code:    common.STUDENT_INVALID_CODE,
				Message: common.STUDENT_INVALID_MESS,
			},
		}, nil
	}
	if topic.GroupStudentsId.Valid {
		groupID = topic.GroupStudentsId.Int64
	}
	_, err = l.svcCtx.StudentGroupModel.Insert(l.ctx, &model.StudentGroupTbl{
		Id:        sync.RandomID(),
		GroupId:   groupID,
		StudentId: req.StudentID,
		EventId:   currentEvent.Id,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicStudentGroupRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	err = l.svcCtx.TopicModel.UpdateGroup(l.ctx, req.ID, groupID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicStudentGroupRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.UpdateTopicStudentGroupRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
