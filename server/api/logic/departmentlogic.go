package logic

// import (
// 	"errors"
// 	db "github/ThoPham02/research_management/api/db/sqlc"
// 	"github/ThoPham02/research_management/api/types"
// )

// func (l *Logic) GetListDepartmentLogic(req *types.GetListDepartmentByFaculityRequest) (*types.GetListDepartmentByFaculityResponse, error) {
// 	if req == nil || req.FaculityID < 0 {
// 		l.logHelper.Errorf("request invalid")
// 		return nil, errors.New("request invalid")
// 	}
// 	l.logHelper.Infof("Start process get list department by faculity ID, faculity id: %v", req.FaculityID)
// 	var data []db.Department
// 	var err error

// 	if req.FaculityID == 0 {
// 		data, err = l.svcCtx.Store.GetListDepartment(l.ctx)
// 		if err != nil {
// 			l.logHelper.Errorf("Failed while getting list department, error: %v", err)
// 			return nil, err
// 		}
// 	} else {
// 		data, err = l.svcCtx.Store.GetListDepartmentByFaculity(l.ctx, req.FaculityID)
// 		if err != nil {
// 			l.logHelper.Errorf("Failed while getting list department, error: %v", err)
// 			return nil, err
// 		}
// 	}

// 	var listDepartment []*types.Department
// 	for _, tmp := range data {
// 		listDepartment = append(listDepartment, &types.Department{
// 			ID:         tmp.ID,
// 			Name:       tmp.Name,
// 			FaculityID: tmp.FaculityID,
// 		})
// 	}
// 	return &types.GetListDepartmentByFaculityResponse{
// 		ListDepartment: listDepartment,
// 	}, nil
// }

// func (l *Logic) GetDepartmentByIDLogic(req *types.GetDepartmentByIDRequest) (*types.GetDepartmentByIDResponse, error) {
// 	if req == nil || req.ID <= 0 {
// 		l.logHelper.Errorf("request invalid")
// 		return nil, errors.New("request invalid")
// 	}

// 	l.logHelper.Infof("Start process get list department by faculity ID, faculity id: %v", req.ID)
// 	data, err := l.svcCtx.Store.GetDepartmentByID(l.ctx, req.ID)
// 	if err != nil {
// 		l.logHelper.Errorf("Failed while getting list department, error: %v", err)
// 		return nil, err
// 	}

// 	return &types.GetDepartmentByIDResponse{
// 		ID:         data.ID,
// 		Name:       data.Name,
// 		FaculityID: data.FaculityID,
// 	}, nil
// }
