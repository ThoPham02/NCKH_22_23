package logic

import (
	"database/sql"
	"errors"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/types"
)

func (l *Logic) GetFaculityByIDLogic(id int32, req *types.GetFacultyByIDRequest) (*types.GetFacultyByIDResponse, error) {
	l.logHelper.Info(req)
	var err error
	if req == nil || id <= 0 {
		err = errors.New("request invalid")
		l.logHelper.Error(err)
		return nil, err
	}

	faculity, err := l.svcCtx.Store.GetFaculty(l.ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetFacultyByIDResponse{
				Result: types.Result{
					Code:    constant.SUCCESS_CODE,
					Message: constant.SUCCESS_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.GetFacultyByIDResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	return &types.GetFacultyByIDResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Faculty: types.Faculty{
			ID:   faculity.ID,
			Name: faculity.Name,
		},
	}, nil
}

func (l *Logic) GetListFaculityLogic(req *types.GetFacultysRequest) (resp *types.GetFacultysResponse, err error) {
	l.logHelper.Info("GetListFaculityLogic ", req)
	var data []types.Faculty

	listFaculty, err := l.svcCtx.Store.ListFaculties(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetFacultysResponse{
				Result: types.Result{
					Code:    constant.SUCCESS_CODE,
					Message: constant.SUCCESS_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.GetFacultysResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	for _, tmp := range listFaculty {

		data = append(data, types.Faculty{
			ID:   tmp.ID,
			Name: tmp.Name,
		})
	}

	return &types.GetFacultysResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Faculties: data,
	}, nil
}
