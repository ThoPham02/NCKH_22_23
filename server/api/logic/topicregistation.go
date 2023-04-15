package logic

import (
	"database/sql"
	"fmt"
	"github/ThoPham02/research_management/api/constant"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
)

func (l *Logic) GetListTopicRegistationLogic(req *types.GetTopicRegistrationsRequest) (resp *types.GetTopicRegistrationsResponse, err error) {
	l.logHelper.Info("GetListTopicRegistationLogic", req)
	var data []types.TopicRegistration
	mapFaculties := make(map[int32]string, 0)
	mapUserInfo := make(map[int32]string, 0)

	list, err := l.svcCtx.Store.ListTopicRegistrations(l.ctx, fmt.Sprintf("%%%s%%", req.Search))
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetTopicRegistrationsResponse{
				Result: types.Result{
					Code:    constant.SUCCESS_CODE,
					Message: constant.SUCCESS_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.GetTopicRegistrationsResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	faculties, err := l.svcCtx.Store.ListFaculties(l.ctx)
	if err != nil {
		l.logHelper.Error(err)
		return &types.GetTopicRegistrationsResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	for _, faculty := range faculties {
		mapFaculties[faculty.ID] = faculty.Name
	}

	users, err := l.svcCtx.Store.ListUserInfosByType(l.ctx, constant.LectureType)
	if err != nil {
		l.logHelper.Error(err)
		return &types.GetTopicRegistrationsResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	for _, user := range users {
		mapUserInfo[user.ID] = user.Name
	}

	for _, tmp := range list {
		if req.Status != 0 && req.Status != tmp.Status {
			continue
		}
		if req.FacultyID != 0 && req.FacultyID != tmp.FacultyID {
			continue
		}
		if req.LectureID != 0 && req.LectureID != tmp.LectureID {
			continue
		}
		data = append(data, types.TopicRegistration{
			ID:      tmp.ID,
			Name:    tmp.Name,
			Lecture: mapUserInfo[tmp.LectureID],
			Faculty: mapFaculties[tmp.FacultyID],
			Status:  constant.MapStatusTopicRegistration[tmp.Status],
		})
	}

	return &types.GetTopicRegistrationsResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		TopicRegistrations: data,
	}, nil
}

func (l *Logic) CreateTopicRegistationLogic(req *types.CreateTopicRegistrationRequest) (resp *types.CreateTopicRegistrationResponse, err error) {
	l.logHelper.Info("CreateTopicRegistationLogic ", req)
	_, err = l.svcCtx.Store.CreateTopicRegistration(l.ctx, db.CreateTopicRegistrationParams{
		ID:        utils.RandomID(),
		Name:      req.Name,
		LectureID: req.LectureID,
		FacultyID: req.FacultyID,
		Status:    1,
	})
	if err != nil {
		l.logHelper.Error(err)
		return &types.CreateTopicRegistrationResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	return &types.CreateTopicRegistrationResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}

func (l *Logic) UpdateTopicRegistationLogic(id int32, req *types.UpdateTopicRegistrationRequest) (resp *types.UpdateTopicRegistrationResponse, err error) {
	l.logHelper.Info("UpdateTopicRegistationLogic", id, req)

	topicRegis, err := l.svcCtx.Store.GetTopicRegistration(l.ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.UpdateTopicRegistrationResponse{
				Result: types.Result{
					Code:    constant.TOPIC_REGIS_NOT_FOUND_CODE,
					Message: constant.TOPIC_REGIS_NOT_FOUND_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.UpdateTopicRegistrationResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	err = l.svcCtx.Store.UpdateTopicRegistration(l.ctx, db.UpdateTopicRegistrationParams{
		ID:        id,
		Name:      topicRegis.Name,
		LectureID: topicRegis.LectureID,
		FacultyID: topicRegis.FacultyID,
		Status:    req.Status,
	})
	if err != nil {
		l.logHelper.Error(err)
		return &types.UpdateTopicRegistrationResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	return &types.UpdateTopicRegistrationResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}

func (l *Logic) GetTopicRegistationByIdLogic(id int32, req *types.GetTopicRegistrationByIdRequest) (resp *types.GetTopicRegistrationByIdResponse, err error) {
	l.logHelper.Info("GetTopicRegistationByIdLogic ", id, req)

	topicRegis, err := l.svcCtx.Store.GetTopicRegistration(l.ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetTopicRegistrationByIdResponse{
				Result: types.Result{
					Code:    constant.TOPIC_REGIS_NOT_FOUND_CODE,
					Message: constant.TOPIC_REGIS_NOT_FOUND_MESSAGE,
				},
			}, nil
		}
		l.logHelper.Error(err)
		return &types.GetTopicRegistrationByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	faculty, err := l.svcCtx.Store.GetFaculty(l.ctx, topicRegis.FacultyID)
	if err != nil {
		l.logHelper.Error(err)
		return &types.GetTopicRegistrationByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	lecture, err := l.svcCtx.Store.GetUserInfo(l.ctx, topicRegis.LectureID)
	if err != nil {
		l.logHelper.Error(err)
		return &types.GetTopicRegistrationByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	return &types.GetTopicRegistrationByIdResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		TopicRegistration: types.TopicRegistration{
			ID:      topicRegis.ID,
			Name:    topicRegis.Name,
			Lecture: lecture.Name,
			Faculty: faculty.Name,
			Status:  constant.MapStatusTopicRegistration[topicRegis.Status],
		},
	}, nil
}
