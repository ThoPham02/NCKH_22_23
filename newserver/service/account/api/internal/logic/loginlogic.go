package logic

import (
	"context"
	"time"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/account/api/internal/svc"
	"github.com/ThoPham02/research_management/service/account/api/internal/types"
	"github.com/ThoPham02/research_management/service/account/api/utils"
	"github.com/ThoPham02/research_management/service/account/model"
	"github.com/golang-jwt/jwt/v4"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// todo: add your logic here and delete this line
	var user *model.UserTbl
	var accessSecret string = l.svcCtx.Config.Auth.AccessSecret
	var accessExpire int64 = l.svcCtx.Config.Auth.AccessExpire
	var refreshSecret string = l.svcCtx.Config.Auth.RefreshSecret
	var refreshExpire int64 = l.svcCtx.Config.Auth.RefreshExpire

	user, err = l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if isCorrect := utils.ComparePassword(req.Password, user.HashPassword); !isCorrect {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.LOGIN_ERR_CODE,
				Message: common.LOGIn_ERR_MESS,
			},
		}, nil
	}

	now := time.Now().Unix()
	accessToken, err := l.getJwtToken(accessSecret, now, accessExpire, user.Id)
	if err != nil {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	refreshToken, err := l.getJwtToken(refreshSecret, now, refreshExpire, user.Id)
	if err != nil {
		return &types.LoginRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	return &types.LoginRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Token: types.Token{
			RefreshToken: accessToken,
			AccessToken:  refreshToken,
		},
		UserInfo: types.User{
			ID:           user.Id,
			Username:     user.Username,
			Role:         user.Role,
			Name:         user.Name,
			Email:        user.Email.String,
			Phone:        user.Phone.String,
			FacultyID:    user.FacultyId,
			DepartmentID: user.Department.Int64,
			YearStart:    user.YearStart,
			Degree:       user.Degree,
			AvatarUrl:    user.AvataUrl.String,
			Birthday:     user.Birthday.String,
			BankAccount:  user.BankAccount.String,
		},
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat int64, seconds int64, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[common.JWT_IAT_KEY] = iat
	claims[common.JWT_EXP_KEY] = iat + seconds
	claims[common.JWT_USER_ID_KEY] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
