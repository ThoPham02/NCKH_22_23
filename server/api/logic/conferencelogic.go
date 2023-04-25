package logic

import (
	"github/ThoPham02/research_management/api/constant"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
)

func (l *Logic) CreateConferenceLogic(req *types.CreateConferenceRequest) (resp *types.CreateConferenceResponse, err error) {
	l.logHelper.Info("CreateConferenceLogic", req)

	_, err = l.svcCtx.Store.CreateConference(l.ctx, db.CreateConferenceParams{
		ID:          utils.RandomID(),
		Name:        req.Name,
		CashSupport: utils.GetInt32(req.CrashSupport),
		SchoolYear:  utils.GetString(req.SchoolYear),
	})
	if err != nil {
		return &types.CreateConferenceResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	return &types.CreateConferenceResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}
