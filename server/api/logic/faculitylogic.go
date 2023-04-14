package logic

// import (
// 	"errors"
// 	"github/ThoPham02/research_management/api/types"
// )

// func (l *Logic) GetFaculityByIDLogic(req *types.GetFaculityByIDRequest) (*types.GetFaculityByIDResponse, error) {
// 	l.logHelper.Info(req)
// 	if req == nil || req.ID <= 0 {
// 		l.logHelper.Errorf("request invalid")
// 		return nil, errors.New("request invalid")
// 	}
// 	l.logHelper.Infof("Start process get faculity by ID, id: %v", req.ID)

// 	faculity, err := l.svcCtx.Store.GetFaculityByID(l.ctx, req.ID)
// 	if err != nil {
// 		l.logHelper.Errorf("Failed while getting facility by id, err: %v", err)
// 		return nil, err
// 	}

// 	return &types.GetFaculityByIDResponse{
// 		ID:   faculity.ID,
// 		Name: faculity.Name,
// 	}, nil
// }

// func (l *Logic) GetListFaculityLogic() (*types.GetListFaculityResponse, error) {
// 	l.logHelper.Infof("Start process get list facility")
// 	data, err := l.svcCtx.Store.GetListFaculity(l.ctx)
// 	if err != nil {
// 		l.logHelper.Errorf("Failed while getting list facility, err: %v", err)
// 		return nil, err
// 	}

// 	var listFaulity []*types.Faculity

// 	for _, tmp := range data {
// 		listFaulity = append(listFaulity, &types.Faculity{
// 			ID:   tmp.ID,
// 			Name: tmp.Name,
// 		})
// 	}

// 	return &types.GetListFaculityResponse{
// 		ListFaculity: listFaulity,
// 	}, nil
// }
