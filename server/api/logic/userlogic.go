package logic

import (
	"github/ThoPham02/research_management/api/types"
)

func (l *Logic) Login(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	// l.logHelper.Infof("Start login process, input: %v", req)
	// user, err := l.svcCtx.Store.GetUserByName(l.ctx, req.Name)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, errors.New(constant.UserIsNotExistErrMsg)
	// 	}
	// 	l.logHelper.Errorf("failed to get user, error: %s", err.Error())
	// 	return nil, err
	// }

	// correctPassword := utils.ComparePassword(req.Password, user.HashPassword)
	// if !correctPassword {
	// 	return nil, errors.New(constant.WrongPasswordErrMsg)
	// }

	// tokenMaker, err := token.NewPasetoMaker(l.svcCtx.Config.TokenSemmetricKey)
	// if err != nil {
	// 	l.logHelper.Errorf("Failed while create token maker, error: %s", err.Error())
	// 	return nil, err
	// }

	// accessToken, err := tokenMaker.CreateToken(user.ID, user.TypeAccountID, l.svcCtx.Config.AccessTokenDuration)
	// if err != nil {
	// 	l.logHelper.Errorf("Failed while create access token, error: %s", err.Error())
	// 	return nil, err
	// }

	// refreshToken, err := tokenMaker.CreateToken(user.ID, user.TypeAccountID, l.svcCtx.Config.RefreshTokenDuration)
	// if err != nil {
	// 	l.logHelper.Errorf("Failed while create refresh token, error: %s", err.Error())
	// 	return nil, err
	// }

	// currency := time.Now()

	// return &types.UserLoginResponse{
	// 	AccessToken:      accessToken,
	// 	AccessExpiredAt:  currency.Add(l.svcCtx.Config.AccessTokenDuration).String(),
	// 	RefreshToken:     refreshToken,
	// 	RefreshExpiredAt: currency.Add(l.svcCtx.Config.RefreshTokenDuration).String(),
	// 	User: types.User{
	// 		Name:        user.Name,
	// 		Email:       user.Email,
	// 		AccountType: user.TypeAccountID,
	// 	},
	// }, nil
	return
}

func (l *Logic) RefreshToken(req *types.UserRefreshTokenRequest) (resp *types.UserRefreshTokenResponse, err error) {
	// l.logHelper.Infof("Start processing refresh token, input: %v", req)
	// tokenMaker, err := token.NewPasetoMaker(l.svcCtx.Config.TokenSemmetricKey)
	// if err != nil {
	// 	l.logHelper.Errorf("Failed while creating token maker, error: %v", err)
	// 	return nil, err
	// }

	// payload, err := tokenMaker.VerifyToken(req.RefreshToken)
	// if err != nil {
	// 	l.logHelper.Errorf("Failed while verify token, error: %v", err)
	// 	return nil, err
	// }

	// accessToken, err := tokenMaker.CreateToken(payload.UserID, payload.AccountType, l.svcCtx.Config.AccessTokenDuration)
	// if err != nil {
	// 	l.logHelper.Errorf("Failed while creating access token, error: %v", err)
	// 	return nil, err
	// }

	// return &types.AccessTokenResponse{
	// 	AccessToken: accessToken,
	// }, nil
	return
}

func (l *Logic) Register(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	return
}

func (l *Logic) UpdateUserInfo(userID int64, req *types.UpdateUserInfoRequest) (resp *types.UpdateUserInfoResponse, err error) {

	return
}

func (l *Logic) GetUserInfo(id int32, req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	return
}
