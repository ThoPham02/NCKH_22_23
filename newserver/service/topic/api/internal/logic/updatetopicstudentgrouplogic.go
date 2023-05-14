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
	var studentGroup []model.StudentGroupTbl

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

	for _, student := range req.ListStudentID {
		studentGroup = append(studentGroup, model.StudentGroupTbl{
			Id:        sync.RandomID(),
			EventId:   currentEvent.Id,
			StudentId: student,
			GroupId:   groupID,
		})
	}

	err = l.svcCtx.StudentGroupModel.InsertMutil(l.ctx, studentGroup)
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
