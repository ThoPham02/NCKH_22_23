package logic

import (
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/api/types"
)

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
