package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetDepartmentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepartmentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentsLogic {
	return &GetDepartmentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepartmentsLogic) GetDepartments(req *types.GetDepartmentsReq) (resp *types.GetDepartmentsRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetDepartments", req)

	var departments []types.Department
	var department types.Department

	var departmentsModel []model.DepartmentTbl
	var departmentModel model.DepartmentTbl
	departmentsModel, err = l.svcCtx.DepartmentModel.FindDepartments(l.ctx, req.FacultyID)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return &types.GetDepartmentsRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetDepartmentsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, departmentModel = range departmentsModel {
		department = types.Department{
			ID:        departmentModel.Id,
			Name:      departmentModel.Name,
			FacultyID: departmentModel.FacultyId,
		}
		departments = append(departments, department)
	}

	return &types.GetDepartmentsRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Departments: departments,
	}, nil
}
