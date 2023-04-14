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

func (l *Logic) UpdateUserInfo(userID int64, req *types.UserInfoResponse) error {
	l.logHelper.Infof("Start processing update user info, input: %d, %v", userID, req)

	description := sql.NullString{
		Valid:  req.Description != nil,
		String: getString(req.Description),
	}
	avataUrl := sql.NullString{
		Valid:  req.AvatarUrl != nil,
		String: getString(req.AvatarUrl),
	}
	birthday := sql.NullTime{
		Valid: req.Birthday != nil,
		Time:  getTime(req.Birthday),
	}
	bankAccount := sql.NullString{
		Valid:  req.BankAccount != nil,
		String: getString(req.BankAccount),
	}
	phone := sql.NullString{
		Valid:  req.Phone != nil,
		String: getString(req.Phone),
	}
	sex := sql.NullInt64{
		Valid: req.Sex != nil,
		Int64: getInt64(req.Sex),
	}

	_, err := l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err == sql.ErrNoRows {
		err = l.svcCtx.Store.CreateUserInfo(l.ctx, db.CreateUserInfoParams{
			UserID:      userID,
			Name:        req.Name,
			Description: description,
			AvataUrl:    avataUrl,
			FaculityID:  req.FaculityID,
			YearStart:   req.YearStart,
			Birthday:    birthday,
			BankAccount: bankAccount,
			Phone:       phone,
			Sex:         sex,
		})
	}
	if err != nil {
		l.logHelper.Errorf("Failed while creating user info, error: %v", err)
		return err
	}

	if err = l.svcCtx.Store.UpdateUserInfo(l.ctx, db.UpdateUserInfoParams{
		UserID:      userID,
		Name:        req.Name,
		Description: description,
		AvataUrl:    avataUrl,
		FaculityID:  req.FaculityID,
		YearStart:   req.YearStart,
		Birthday:    birthday,
		BankAccount: bankAccount,
		Phone:       phone,
		Sex:         sex,
	}); err != nil {
		l.logHelper.Errorf("Failed while update user info: %v", err)
		return err
	}

	return nil
}

func getString(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func getInt64(input *int64) int64 {
	if input == nil {
		return 0
	}
	return *input
}

func getTime(input *time.Time) time.Time {
	if input == nil {
		return time.Now()
	}
	return *input
}

func (l *Logic) GetUserInfo(userID int64) (*types.UserInfoResponse, error) {
	l.logHelper.Infof("Start processing get user info, input: %d", userID)

	userInfo, err := l.svcCtx.Store.GetUserInfo(l.ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.UserInfoResponse{}, nil
		}
		l.logHelper.Errorf("Failed to get user info: %s", err.Error())
		return nil, err
	}

	var (
		description *string
		avataUrl    *string
		birthday    *time.Time
		bankAccount *string
		phone       *string
		sex         *int64
	)
	if userInfo.Description.Valid {
		description = &userInfo.Description.String
	}
	if userInfo.Birthday.Valid {
		birthday = &userInfo.Birthday.Time
	}
	if userInfo.AvataUrl.Valid {
		avataUrl = &userInfo.AvataUrl.String
	}
	if userInfo.BankAccount.Valid {
		bankAccount = &userInfo.BankAccount.String
	}
	if userInfo.Phone.Valid {
		phone = &userInfo.Phone.String
	}
	if userInfo.Sex.Valid {
		sex = &userInfo.Sex.Int64
	}

	return &types.UserInfoResponse{
		Name:        userInfo.Name,
		Description: description,
		AvatarUrl:   avataUrl,
		Birthday:    birthday,
		FaculityID:  userInfo.FaculityID,
		YearStart:   userInfo.YearStart,
		BankAccount: bankAccount,
		Phone:       phone,
		Sex:         sex,
	}, nil
}
