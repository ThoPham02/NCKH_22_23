package logic

import (
	"database/sql"
	"errors"
	"github/ThoPham02/research_management/api/constant"
	db "github/ThoPham02/research_management/api/db/sqlc"
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
		AccessExpiredAt:  currency.Add(l.svcCtx.Config.AccessTokenDuration).String(),
		RefreshToken:     refreshToken,
		RefreshExpiredAt: currency.Add(l.svcCtx.Config.RefreshTokenDuration).String(),
		User: types.User{
			Name:        user.Name,
			Email:       user.Email,
			AccountType: user.TypeAccountID,
		},
	}, nil
}

func (l *Logic) RefreshToken(req *types.RefreshTokenRequest) (*types.AccessTokenResponse, error) {
	l.logHelper.Infof("Start processing refresh token, input: %v", req)
	tokenMaker, err := token.NewPasetoMaker(l.svcCtx.Config.TokenSemmetricKey)
	if err != nil {
		l.logHelper.Errorf("Failed while creating token maker, error: %v", err)
		return nil, err
	}

	payload, err := tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		l.logHelper.Errorf("Failed while verify token, error: %v", err)
		return nil, err
	}

	accessToken, err := tokenMaker.CreateToken(payload.UserID, payload.AccountType, l.svcCtx.Config.AccessTokenDuration)
	if err != nil {
		l.logHelper.Errorf("Failed while creating access token, error: %v", err)
		return nil, err
	}

	return &types.AccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}

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

func (l *Logic) UpdateUserInfo(userID int64, req *types.UpdateUserInfoRequest) (*types.UpdateUserInfoResponse, error) {
	l.logHelper.Infof("Start processing update user info, input: %d, %v", userID, req)

	var err error
	var birthday *time.Time

	if req.Birthday != nil {
		birthday, err = utils.ConvertStringToTime(*req.Birthday)
		if err != nil {
			l.logHelper.Error(err)
			return nil, err
		}
	}

	_, err = l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err == sql.ErrNoRows {
		err = l.svcCtx.Store.CreateUserInfo(l.ctx, db.CreateUserInfoParams{
			UserID:      userID,
			Name:        req.Name,
			Description: utils.GetString(req.Description),
			AvataUrl:    utils.GetString(req.AvatarUrl),
			FaculityID:  req.FaculityID,
			YearStart:   req.YearStart,
			Birthday:    utils.GetTime(birthday),
			BankAccount: utils.GetString(req.BankAccount),
			Phone:       utils.GetString(req.Phone),
			Sex:         utils.GetInt64(req.Sex),
		})
	}
	if err != nil {
		l.logHelper.Errorf("Failed while creating user info, error: %v", err)
		return nil, err
	}

	if err = l.svcCtx.Store.UpdateUserInfo(l.ctx, db.UpdateUserInfoParams{
		UserID:      userID,
		Name:        req.Name,
		Description: utils.GetString(req.Description),
		AvataUrl:    utils.GetString(req.AvatarUrl),
		FaculityID:  req.FaculityID,
		YearStart:   req.YearStart,
		Birthday:    utils.GetTime(birthday),
		BankAccount: utils.GetString(req.BankAccount),
		Phone:       utils.GetString(req.Phone),
		Sex:         utils.GetInt64(req.Sex),
	}); err != nil {
		l.logHelper.Errorf("Failed while update user info: %v", err)
		return nil, err
	}
	return &types.UpdateUserInfoResponse{}, nil
}

func (l *Logic) GetUserInfo(userID int64) (*types.GetUserInfoResponse, error) {
	l.logHelper.Infof("Start processing get user info, input: %d", userID)

	userInfo, err := l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetUserInfoResponse{}, nil
		}
		l.logHelper.Errorf("Failed to get user info: %s", err.Error())
		return nil, err
	}

	return &types.GetUserInfoResponse{
		Name:        userInfo.Name,
		Description: userInfo.Description.String,
		AvatarUrl:   userInfo.AvataUrl.String,
		Birthday:    userInfo.Birthday.Time.String(),
		FaculityID:  userInfo.FaculityID,
		YearStart:   userInfo.YearStart,
		BankAccount: userInfo.BankAccount.String,
		Phone:       userInfo.Phone.String,
		Sex:         userInfo.Sex.Int64,
	}, nil
}
