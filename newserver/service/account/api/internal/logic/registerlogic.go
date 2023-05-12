package logic

import (
	"context"
	"database/sql"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/account/api/internal/svc"
	"github.com/ThoPham02/research_management/service/account/api/internal/types"
	"github.com/ThoPham02/research_management/service/account/api/utils"
	"github.com/ThoPham02/research_management/service/account/model"
	"github.com/ThoPham02/research_management/sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	// todo: add your logic here and delete this line
	var hashPassword string
	hashPassword, err = utils.HashPassword(req.Password)
	if err != nil {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.INVALID_INPUT_ERR_CODE,
				Message: common.INVALID_INPUT_ERR_MESS,
			},
		}, nil
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.UserTbl{
		Id:           sync.RandomID(),
		Username:     req.Username,
		HashPassword: hashPassword,
		Role:         req.Role,
		Name:         req.Name,
		Email: sql.NullString{
			Valid:  req.Email != "",
			String: req.Email,
		},
		Phone: sql.NullString{
			Valid:  req.Phone != "",
			String: req.Phone,
		},
		FacultyId: req.FacultyID,
		YearStart: req.YearStart,
		Degree:    req.Degree,
		AvataUrl: sql.NullString{
			Valid:  req.AvatarUrl != "",
			String: req.AvatarUrl},
		Birthday: sql.NullString{
			Valid:  req.Birthday != "",
			String: req.Birthday,
		},
		BankAccount: sql.NullString{
			Valid:  req.BankAccount != "",
			String: req.BankAccount,
		},
	})
	if err != nil {
		return &types.RegisterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.RegisterRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
