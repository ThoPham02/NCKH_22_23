package logic

import (
	"database/sql"
	"errors"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
	"time"
)

func (l *Logic) Login(req *types.UserLoginRequest) (*types.UserLoginResponse, error) {
	l.logHelper.Infof("Start login process, input: %v", req)
	user, err := l.svcCtx.Store.GetUserByName(l.ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.UserIsNotExistErrMsg)
		}
		l.logHelper.Errorf("failed to get user, error: %s", err.Error())
		return nil, err
	}

	correctPassword := utils.ComparePassword(req.Password, user.HashPassword)
	if !correctPassword {
		return nil, errors.New(constant.WrongPasswordErrMsg)
	}

	tokenMaker, err := token.NewPasetoMaker(l.svcCtx.Config.TokenSemmetricKey)
	if err != nil {
		l.logHelper.Errorf("Failed while create token maker, error: %s", err.Error())
		return nil, err
	}

	accessToken, err := tokenMaker.CreateToken(user.ID, user.TypeAccountID, l.svcCtx.Config.AccessTokenDuration)
	if err != nil {
		l.logHelper.Errorf("Failed while create access token, error: %s", err.Error())
		return nil, err
	}

	refreshToken, err := tokenMaker.CreateToken(user.ID, user.TypeAccountID, l.svcCtx.Config.RefreshTokenDuration)
	if err != nil {
		l.logHelper.Errorf("Failed while create refresh token, error: %s", err.Error())
		return nil, err
	}

	currency := time.Now()

	return &types.UserLoginResponse{
		AccessToken:      accessToken,
		AccessExpiredAt:  currency.Add(l.svcCtx.Config.AccessTokenDuration),
		RefreshToken:     refreshToken,
		RefreshExpiredAt: currency.Add(l.svcCtx.Config.RefreshTokenDuration),
		User: types.User{
			Name:        user.Name,
			Email:       user.Email,
			AccountType: user.TypeAccountID,
		},
	}, nil
}
