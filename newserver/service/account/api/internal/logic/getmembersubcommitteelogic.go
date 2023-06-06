package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/account/api/internal/svc"
	"github.com/ThoPham02/research_management/service/account/api/internal/types"
	"github.com/ThoPham02/research_management/service/account/model"
	resultModel "github.com/ThoPham02/research_management/service/result/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberSubcommitteeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberSubcommitteeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberSubcommitteeLogic {
	return &GetMemberSubcommitteeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberSubcommitteeLogic) GetMemberSubcommittee(req *types.MemberSubcommitteesReq) (resp *types.MemberSubcommitteesRes, err error) {
	// todo: add your logic here and delete this line
	var listSub []int64
	var listLecture []int64
	var memberSubcommittee types.MemberSubcommittee
	var memberSubcommittees []types.MemberSubcommittee

	l.Logger.Info(req)

	event, err := l.svcCtx.EventModel.FindCurrentEvent(l.ctx)
	if err != nil || event == nil {
		l.Logger.Error(err)
		return &types.MemberSubcommitteesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	subcommittees, err := l.svcCtx.SubcommitteeModel.FindSubcommittee(l.ctx, resultModel.SubcommitteeConditions{EventID: event.Id})
	if err != nil {
		l.Logger.Error(err)
		return &types.MemberSubcommitteesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, tmp := range subcommittees {
		listSub = append(listSub, tmp.Id)
	}

	groups, err := l.svcCtx.GroupModel.FindMultiBySubcommittee(l.ctx, listSub)
	if err != nil {
		l.Logger.Error(err)
		return &types.MemberSubcommitteesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, tmp := range groups {
		listLecture = append(listLecture, tmp.LectureId)
	}

	users, err := l.svcCtx.UserModel.FindUserByCondition(l.ctx, model.UserCondition{
		Name:      req.Name,
		Role:      req.Role,
		FacultyID: req.FacultyID,
		NotUser:   listLecture,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.MemberSubcommitteesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, tmp := range users {
		memberSubcommittee = types.MemberSubcommittee{
			ID:     tmp.Id,
			Name:   tmp.Name,
			Degree: tmp.Degree,
		}
		memberSubcommittees = append(memberSubcommittees, memberSubcommittee)
	}

	return &types.MemberSubcommitteesRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Members: memberSubcommittees,
	}, nil
}
