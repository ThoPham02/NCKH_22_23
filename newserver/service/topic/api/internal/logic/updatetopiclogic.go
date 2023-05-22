package logic

import (
	"context"
	"database/sql"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicLogic) UpdateTopic(req *types.UpdateTopicReq) (resp *types.UpdateTopicRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("UpdateTopicStatus", req)

	_, err = l.svcCtx.TopicModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UpdateTopicRes{
				Result: types.Result{
					Code:    common.TOPIC_NOT_EXIST_CODE,
					Message: common.TOPIC_NOT_EXIST_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.UpdateTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.TopicModel.Update(l.ctx, &model.TopicTbl{
		Id:           req.ID,
		Name:         req.Name,
		LectureId:    req.LectureID,
		DepartmentId: req.DepartmentID,
		Status:       req.Status,
		EventId:      req.EventID,
		SubcommitteeId: sql.NullInt64{
			Valid: req.SubcommitteeID != 0,
			Int64: req.SubcommitteeID,
		},
		TimeStart: sql.NullInt64{
			Valid: req.TimeStart != 0,
			Int64: req.TimeStart,
		},
		TimeEnd: sql.NullInt64{
			Valid: req.TimeEnd != 0,
			Int64: req.TimeEnd,
		},
		CashSupport: sql.NullInt64{
			Valid: req.CashSupport != 0,
			Int64: req.CashSupport,
		},
		GroupStudentsId: sql.NullInt64{
			Valid: req.GroupID != 0,
			Int64: req.GroupID,
		},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.UpdateTopicRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
