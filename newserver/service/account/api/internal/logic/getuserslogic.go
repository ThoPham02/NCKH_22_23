package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/account/api/internal/svc"
	"github.com/ThoPham02/research_management/service/account/api/internal/types"
	"github.com/ThoPham02/research_management/service/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req *types.GetUsersReq) (resp *types.GetUsersRes, err error) {
	// todo: add your logic here and delete this line
	var usersModel []model.UserTbl
	var userModel model.UserTbl
	var users []types.User
	var user types.User
	usersModel, err = l.svcCtx.UserModel.FindUserByCondition(l.ctx, model.UserCondtion{})
	if err != nil {
		l.Logger.Error(err)
		return &types.GetUsersRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, userModel = range usersModel {
		user = types.User{
			ID:          userModel.Id,
			Username:    userModel.Username,
			Role:        userModel.Role,
			Name:        userModel.Name,
			Email:       userModel.Email.String,
			Phone:       userModel.Phone.String,
			FacultyID:   userModel.FacultyId,
			YearStart:   userModel.YearStart,
			Degree:      userModel.Degree,
			AvatarUrl:   userModel.AvataUrl.String,
			Birthday:    userModel.Birthday.String,
			BankAccount: userModel.BankAccount.String,
		}
		users = append(users, user)
	}

	return &types.GetUsersRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		// Total: ,
		Users: users,
	}, nil
}
