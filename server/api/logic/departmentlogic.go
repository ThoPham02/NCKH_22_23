package logic

import (
	"database/sql"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/types"
)

func (l *Logic) GetListDepartmentLogic(req *types.GetDepartmentsRequest) (resp *types.GetDepartmentsResponse, err error) {
	l.logHelper.Info("GetListDepartmentLogic ", req)

	var data []types.Department

	listDepartment, err := l.svcCtx.Store.ListDepartments(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetDepartmentsResponse{
				Result: types.Result{
					Code:    constant.SUCCESS_CODE,
					Message: constant.SUCCESS_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.GetDepartmentsResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	for _, tmp := range listDepartment {
		if req.FacultyID != 0 && req.FacultyID == tmp.FacultyID {
			data = append(data, types.Department{
				ID:        tmp.ID,
				Name:      tmp.Name,
				FacultyID: tmp.FacultyID,
			})
		}
	}

	return &types.GetDepartmentsResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Departments: data,
	}, nil
}

func (l *Logic) GetDepartmentByIDLogic(id int32, req *types.GetDepartmentByIDRequest) (resp *types.GetDepartmentByIDResponse, err error) {
	l.logHelper.Info("GetDepartmentByIDLogic ", id, req)

	department, err := l.svcCtx.Store.GetDepartment(l.ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetDepartmentByIDResponse{
				Result: types.Result{
					Code:    constant.SUCCESS_CODE,
					Message: constant.SUCCESS_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.GetDepartmentByIDResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	return &types.GetDepartmentByIDResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Department: types.Department{
			ID:        department.ID,
			Name:      department.Name,
			FacultyID: department.FacultyID,
		},
	}, nil
}
