package logic

import (
	"database/sql"
	"github/ThoPham02/research_management/api/constant"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
	"time"
)

func (l *Logic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	l.logHelper.Info("UserLogin", req)
	user, err := l.svcCtx.Store.GetUserByName(l.ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.UserLoginResponse{
				Result: types.Result{
					Code:    constant.USER_NOT_FOUND_CODE,
					Message: constant.USER_NOT_FOUND_MESSAGE,
				},
			}, nil
		}
		return &types.UserLoginResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	correctPassword := utils.ComparePassword(req.Password, user.HashPassword)
	if !correctPassword {
		return &types.UserLoginResponse{
			Result: types.Result{
				Code:    constant.USER_NOT_FOUND_CODE,
				Message: constant.USER_NOT_FOUND_MESSAGE,
			},
		}, nil
	}

	tokenMaker, err := token.NewPasetoMaker(l.svcCtx.Config.TokenSemmetricKey)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UserLoginResponse{
			Result: types.Result{
				Code:    constant.TOKEN_ERR_CODE,
				Message: constant.TOKEN_ERR_CODE_MESSAGE,
			},
		}, nil
	}

	accessToken, err := tokenMaker.CreateToken(user.ID, user.TypeAccount, l.svcCtx.Config.AccessTokenDuration)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UserLoginResponse{
			Result: types.Result{
				Code:    constant.TOKEN_ERR_CODE,
				Message: constant.TOKEN_ERR_CODE_MESSAGE,
			},
		}, nil
	}

	refreshToken, err := tokenMaker.CreateToken(user.ID, user.TypeAccount, l.svcCtx.Config.RefreshTokenDuration)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UserLoginResponse{
			Result: types.Result{
				Code:    constant.TOKEN_ERR_CODE,
				Message: constant.TOKEN_ERR_CODE_MESSAGE,
			},
		}, nil
	}

	currency := time.Now()
	return &types.UserLoginResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Token: types.Token{
			AccessToken:      accessToken,
			AccessExpiresAt:  currency.Add(l.svcCtx.Config.AccessTokenDuration).Format(time.RFC3339),
			RefreshToken:     refreshToken,
			RefreshExpiresAt: currency.Add(l.svcCtx.Config.RefreshTokenDuration).Format(time.RFC3339),
		},
		User: types.User{
			Name:        user.Name,
			Email:       user.Email,
			TypeAccount: user.TypeAccount,
		},
	}, nil
}

func (l *Logic) RefreshToken(req *types.UserRefreshTokenRequest) (resp *types.UserRefreshTokenResponse, err error) {
	l.logHelper.Info("RefreshToken", req)
	current := time.Now()
	if valid, err := utils.CompareTimeStringWithNow(req.RefreshExpiresAt); valid || err != nil {
		if err != nil {
			l.logHelper.Error(err)
			return &types.UserRefreshTokenResponse{
				Result: types.Result{
					Code:    constant.TOKEN_EXPIRES_CODE,
					Message: constant.TOKEN_EXPIRES_MESSAGE,
				},
			}, nil
		}
	}
	tokenMaker, err := token.NewPasetoMaker(l.svcCtx.Config.TokenSemmetricKey)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UserRefreshTokenResponse{
			Result: types.Result{
				Code:    constant.TOKEN_ERR_CODE,
				Message: constant.TOKEN_ERR_CODE_MESSAGE,
			},
		}, nil
	}

	payload, err := tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UserRefreshTokenResponse{
			Result: types.Result{
				Code:    constant.TOKEN_ERR_CODE,
				Message: constant.TOKEN_ERR_CODE_MESSAGE,
			},
		}, nil
	}

	accessToken, err := tokenMaker.CreateToken(payload.UserID, payload.TypeAccount, l.svcCtx.Config.AccessTokenDuration)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UserRefreshTokenResponse{
			Result: types.Result{
				Code:    constant.TOKEN_ERR_CODE,
				Message: constant.TOKEN_ERR_CODE_MESSAGE,
			},
		}, nil
	}
	return &types.UserRefreshTokenResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Token: types.Token{
			AccessToken:      accessToken,
			AccessExpiresAt:  current.Add(l.svcCtx.Config.AccessTokenDuration).String(),
			RefreshToken:     req.RefreshToken,
			RefreshExpiresAt: req.RefreshExpiresAt,
		},
	}, nil
}

func (l *Logic) Register(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	l.logHelper.Info("Register", req)

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UserRegisterResponse{
			Result: types.Result{
				Code:    constant.HASH_PASSWORD_ERR_CODE,
				Message: constant.HASH_PASSWORD_ERR_MESSAGE,
			},
		}, nil
	}

	user, err := l.svcCtx.Store.CreateUser(l.ctx, db.CreateUserParams{
		ID:           utils.RandomID(),
		Name:         req.Name,
		Email:        req.Email,
		HashPassword: hashPassword,
		TypeAccount:  req.TypeAccount,
	})
	if err != nil {
		return &types.UserRegisterResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	return &types.UserRegisterResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		User: types.User{
			Name:        user.Name,
			Email:       user.Email,
			TypeAccount: user.TypeAccount,
		},
	}, nil
}

func (l *Logic) UpdateUserInfo(userID int32, req *types.UpdateUserInfoRequest) (resp *types.UpdateUserInfoResponse, err error) {
	l.logHelper.Info("UpdateUserInfo", req)
	_, err = l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = l.svcCtx.Store.CreateUserInfo(l.ctx, db.CreateUserInfoParams{
				ID:          utils.RandomID(),
				UserID:      userID,
				Name:        req.Name,
				Email:       req.Email,
				Phone:       req.Phone,
				FacultyID:   req.FacultyID,
				Degree:      req.Degree,
				YearStart:   req.YearStart,
				AvataUrl:    utils.GetString(req.AvatarUrl),
				Birthday:    utils.GetString(req.Birthday),
				BankAccount: utils.GetString(req.BankAccount),
			})
			if err == nil {
				return &types.UpdateUserInfoResponse{
					Result: types.Result{
						Code:    constant.SUCCESS_CODE,
						Message: constant.SUCCESS_MESSAGE,
					},
				}, nil
			}
		}
		l.logHelper.Error(err)
		return &types.UpdateUserInfoResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	err = l.svcCtx.Store.UpdateUserInfo(l.ctx, db.UpdateUserInfoParams{
		UserID:      userID,
		Name:        req.Name,
		Email:       req.Email,
		Phone:       req.Phone,
		FacultyID:   req.FacultyID,
		Degree:      req.Degree,
		YearStart:   req.YearStart,
		AvataUrl:    utils.GetString(req.AvatarUrl),
		Birthday:    utils.GetString(req.Birthday),
		BankAccount: utils.GetString(req.BankAccount),
	})
	if err != nil {
		return &types.UpdateUserInfoResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	return &types.UpdateUserInfoResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}

func (l *Logic) GetUserInfo(userID int32, req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	l.logHelper.Info("GetUserInfo ", userID, req)
	userInfo, err := l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err != nil {
		return &types.GetUserInfoResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	faculty, err := l.svcCtx.Store.GetFaculty(l.ctx, userInfo.FacultyID)
	if err != nil {
		return &types.GetUserInfoResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	return &types.GetUserInfoResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		UserInfo: types.UserInfo{
			ID:          userInfo.ID,
			UserID:      userInfo.UserID,
			Name:        userInfo.Name,
			Email:       userInfo.Email,
			Phone:       userInfo.Phone,
			Faculty:     faculty.Name,
			Degree:      constant.MapDegree[userInfo.Degree],
			YearStart:   userInfo.YearStart,
			AvatarUrl:   userInfo.AvataUrl.String,
			Birthday:    userInfo.Birthday.String,
			BankAccount: userInfo.BankAccount.String,
		},
	}, nil
}
