package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"
	"github.com/ThoPham02/research_management/service/result/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubcommitteeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubcommitteeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubcommitteeLogic {
	return &GetSubcommitteeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubcommitteeLogic) GetSubcommittee(req *types.GetSubcommitteesReq) (resp *types.GetSubcommitteesRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetSubcommittee", req)

	var subcommitteeModel model.SubcommitteeTbl
	var subcommitteesModel []model.SubcommitteeTbl
	var subcommittee types.Subcommittee
	var subcommittees []types.Subcommittee

	subcommitteesModel, err = l.svcCtx.SubcommitteeModel.FindSubcommittee(l.ctx, model.SubcommitteeConditions{})
	if err != nil {
		if err == model.ErrNotFound {
			return &types.GetSubcommitteesRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetSubcommitteesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, subcommitteeModel = range subcommitteesModel {
		subcommittee = types.Subcommittee{
			ID:        subcommitteeModel.Id,
			Name:      subcommitteeModel.Name,
			FacultyID: subcommitteeModel.FacultId.Int64,
			EventID:   subcommitteeModel.EventId,
		}
		subcommittees = append(subcommittees, subcommittee)
	}

	return &types.GetSubcommitteesRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Subcommittees: subcommittees,
		Total:         int64(len(subcommittees)),
	}, nil
}
