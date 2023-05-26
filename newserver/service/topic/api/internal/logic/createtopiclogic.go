package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"
	"github.com/ThoPham02/research_management/sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTopicLogic {
	return &CreateTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTopicLogic) CreateTopic(req *types.CreateTopicReq) (resp *types.CreateTopicRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("CreateTopic", req)

	_, err = l.svcCtx.TopicModel.Insert(l.ctx, &model.TopicTbl{
		Id:           sync.RandomID(),
		Name:         req.Name,
		LectureId:    req.LectureID,
		DepartmentId: req.DepartmentID,
		Status:       common.TOPIC_SUGGESTION,
		EventId:      req.EventID,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.CreateTopicRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
