package logic

import (
	"context"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/api/types"

	"github.com/go-kratos/kratos/v2/log"
)

type UserLoginLogic struct {
	ctx       context.Context
	svc       *service.ServiceContext
	logHelper *log.Helper
}

func NewUserLoginLogic(ctx context.Context, svcContext *service.ServiceContext, logHelper *log.Helper) UserLoginLogic {
	return UserLoginLogic{
		ctx:       ctx,
		svc:       svcContext,
		logHelper: logHelper,
	}
}

func (l *UserLoginLogic) Login(req *types.UserLoginRequest) (*types.UserLoginResponse, error) {
	l.logHelper.Infof("Start login process, input %v", req)

	return &types.UserLoginResponse{}, nil
}
