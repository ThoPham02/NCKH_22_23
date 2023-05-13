package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFacultiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFacultiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFacultiesLogic {
	return &GetFacultiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFacultiesLogic) GetFaculties(req *types.GetFacultiesReq) (resp *types.GetFacultiesRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetFaculties", req)

	var facultiesModel []model.FacultyTbl
	var facultyModel model.FacultyTbl
	var faculties []types.Faculty
	var faculty types.Faculty

	facultiesModel, err = l.svcCtx.FacultyModel.FindFaculties(l.ctx)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.GetFacultiesRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetFacultiesRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, facultyModel = range facultiesModel {
		faculty = types.Faculty{
			ID:   facultyModel.Id,
			Name: facultyModel.Name,
		}
		faculties = append(faculties, faculty)
	}

	return &types.GetFacultiesRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Faculties: faculties,
	}, nil
}
