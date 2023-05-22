package logic

import (
	"context"
	"database/sql"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"
	"github.com/ThoPham02/research_management/service/result/model"
	"github.com/ThoPham02/research_management/sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTopicMarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTopicMarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTopicMarkLogic {
	return &CreateTopicMarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTopicMarkLogic) CreateTopicMark(req *types.CreateTopicMarkReq) (resp *types.CreateTopicMarkRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("CreateTopicMark", req)

	_, err = l.svcCtx.TopicMark.Insert(l.ctx, &model.TopicMarkTbl{
		Id:        sync.RandomID(),
		TopicId:   req.TopicID,
		LectureId: req.LectureID,
		Point:     req.Point,
		Comment:   sql.NullString{Valid: true, String: req.Comment},
		Url:       sql.NullString{Valid: true, String: req.Url},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateTopicMarkRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.CreateTopicMarkRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
