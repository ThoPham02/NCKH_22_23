package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	accountModel "github.com/ThoPham02/research_management/service/account/model"
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
	var mapSubcommitteeGroup = map[int64][]types.Group{}
	var mapLecture = map[int64]types.User{}

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
	lectures, err := l.svcCtx.UserModel.FindUserByCondition(l.ctx, accountModel.UserCondition{
		Name:      "",
		Role:      2,
		FacultyID: 0,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetSubcommitteesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, tmp := range lectures {
		mapLecture[tmp.Id] = types.User{
			ID:           tmp.Id,
			Name:         tmp.Name,
			Email:        tmp.Email.String,
			Phone:        tmp.Phone.String,
			FacultyID:    tmp.FacultyId,
			DepartmentID: tmp.Department.Int64,
			Degree:       tmp.Degree,
		}
	}

	groups, err := l.svcCtx.GroupModel.FindMulti(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetSubcommitteesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, tmp := range groups {
		mapSubcommitteeGroup[tmp.SubcommitteeId] = append(mapSubcommitteeGroup[tmp.SubcommitteeId], types.Group{
			ID:             tmp.Id,
			SubcommitteeId: tmp.SubcommitteeId,
			Lecture:        mapLecture[tmp.LectureId],
			Role:           tmp.Role,
		})
	}

	l.Logger.Info(mapSubcommitteeGroup)
	for _, subcommitteeModel = range subcommitteesModel {
		subcommittee = types.Subcommittee{
			ID:        subcommitteeModel.Id,
			Name:      subcommitteeModel.Name,
			FacultyID: subcommitteeModel.FacultId.Int64,
			EventID:   subcommitteeModel.EventId,
			Groups:    mapSubcommitteeGroup[subcommitteeModel.Id],
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
