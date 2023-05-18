package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/progress/model"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStudentGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentGroupLogic {
	return &DeleteStudentGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStudentGroupLogic) DeleteStudentGroup(req *types.DeleteStudentGroupbyStudentIdReq) (resp *types.DeleteStudentGroupbyStudentIdRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("DeleteStudentGroup", req)

	var event *model.EventTbl

	event, err = l.svcCtx.EventModel.FindCurrentEvent(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteStudentGroupbyStudentIdRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	err = l.svcCtx.StudentGroupModel.DeleteByStudentID(l.ctx, req.StudentID, event.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteStudentGroupbyStudentIdRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.DeleteStudentGroupbyStudentIdRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
