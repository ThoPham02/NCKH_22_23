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

type CreateSubcommitteeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSubcommitteeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubcommitteeLogic {
	return &CreateSubcommitteeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubcommitteeLogic) CreateSubcommittee(req *types.CreateSubcommitteeReq) (resp *types.CreateSubcommitteeRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("CreateSubcommittee", req)
	var subcommitteeID int64 = sync.RandomID()

	var groupModel model.GroupTbl
	var groupsModel []model.GroupTbl
	var level int64 = common.LEVEL_SUBCOMMITTEE

	if req.FacultyID == 0 {
		level = common.LEVEL_SCHOOL
	}

	_, err = l.svcCtx.SubcommitteeModel.Insert(l.ctx, &model.SubcommitteeTbl{
		Id:       subcommitteeID,
		Name:     req.Name,
		FacultId: sql.NullInt64{Valid: true, Int64: req.FacultyID},
		EventId:  req.EventID,
		Level:    level,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateSubcommitteeRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, group := range req.ListLectures {
		groupModel = model.GroupTbl{
			Id:             sync.RandomID(),
			SubcommitteeId: subcommitteeID,
			LectureId:      group.LectureID,
			Role:           group.Role,
		}
		groupsModel = append(groupsModel, groupModel)
	}

	err = l.svcCtx.GroupModel.InsertMutil(l.ctx, groupsModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateSubcommitteeRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.CreateSubcommitteeRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
