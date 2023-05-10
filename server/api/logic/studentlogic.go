package logic

import (
	"fmt"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/types"
)

func (l *Logic) SearchStudentLogic(req *types.SearchStudentRequest) (resp *types.SearchStudentResponse, err error) {
	l.logHelper.Info("Register", req)
	var data []types.StudentInfo

	listUser, err := l.svcCtx.Store.GetStudentByName(l.ctx, fmt.Sprintf("%%%s%%", req.Name))
	if err != nil {
		l.logHelper.Error(err)
		return &types.SearchStudentResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	for _, user := range listUser {
		typeAccount, err := l.svcCtx.Store.GetTypeAccount(l.ctx, user.UserID)
		if err != nil {
			l.logHelper.Error(err)
			return &types.SearchStudentResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}

		if typeAccount != 1 {
			continue
		}

		data = append(data, types.StudentInfo{
			Name: user.Name,
			ID:   user.UserID,
		})
	}

	return &types.SearchStudentResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		StudentList: data,
	}, nil
}
