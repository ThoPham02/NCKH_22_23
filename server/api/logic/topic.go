package logic

import (
	"database/sql"
	"github/ThoPham02/research_management/api/constant"
	"github/ThoPham02/research_management/api/types"
)

func (l *Logic) GetTopicByIdLogic(id int32) (resp *types.GetTopicByIdResponse, err error) {
	l.logHelper.Info("GetTopicByIdLogic: ", id)

	data, err := l.svcCtx.Store.GetTopic(l.ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetTopicByIdResponse{
				Result: types.Result{
					Code:    constant.SUCCESS_CODE,
					Message: constant.SUCCESS_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.GetTopicByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	lecture, err := l.svcCtx.Store.GetUserInfo(l.ctx, data.LectureID)
	if err != nil {
		l.logHelper.Error(err)
		return &types.GetTopicByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	faculty, err := l.svcCtx.Store.GetFaculty(l.ctx, data.FacultyID)
	if err != nil {
		l.logHelper.Error(err)
		return &types.GetTopicByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	var listName = []string{}
	listStudents, err := l.svcCtx.Store.ListStudentByTopicId(l.ctx, data.ID)
	if err != nil {
		return &types.GetTopicByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	l.svcCtx.Store.GetUserNameByListID(l.ctx, listStudents)
	

	return &types.GetTopicByIdResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Topic: types.Topic{
			ID:           id,
			Name:         data.Name,
			Lecture:      lecture.Name,
			Faculty:      faculty.Name,
			Status:       constant.MapStatusTopic[data.Status],
			ResultUrl:    data.ResultUrl.String,
			ListStudents: []string{},
			TimeStart:    data.TimeStart.Format(constant.Layout),
			TimeEnd:      data.TimeEnd.Format(constant.Layout),
		},
	}, nil
}

func (l *Logic) GetTopicLogic(req *types.GetTopicRequest) (resp *types.GetTopicResponse, err error) {
	return
}

func (l *Logic) AcceptTopicLogic(req *types.AcceptTopicRequest) (resp *types.AcceptTopicResponse, err error) {
	return
}

func (l *Logic) CreateTopicLogic(req *types.CreateTopicRequest) (resp *types.CreateTopicResponse, err error) {
	return
}

func (l *Logic) UpdateTopicLogic(id int32, req *types.UpdateTopicRequest) (resp *types.UpdateTopicResponse, err error) {
	return
}
