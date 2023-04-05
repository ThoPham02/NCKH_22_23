package logic

import (
	"errors"
	"github/ThoPham02/research_management/api/constant"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
)

func (l *Logic) Register(req *types.UserRegisterRequest) error {
	l.logHelper.Infof("Start processing regiter user, input: %v", req)

	if isEmail := utils.ValidateEmail(req.Email); !isEmail {
		return errors.New(constant.InputValidationErrMsg)
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		l.logHelper.Errorf("failed to hash password, error: %s", err.Error())
		return err
	}

	err = l.svcCtx.Store.CreateUser(l.ctx, db.CreateUserParams{
		Name:          req.Name,
		HashPassword:  hashPassword,
		Email:         req.Email,
		TypeAccountID: req.TypeAccountID,
	})
	if err != nil {
		l.logHelper.Errorf("failed while creating user: %s", err.Error())
		return err
	}

	return nil
}
