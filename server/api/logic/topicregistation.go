package logic

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github/ThoPham02/research_management/api/constant"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
	"time"
)

func (l *Logic) GetListTopicRegistationLogic(req *types.GetTopicRegistrationsRequest) (resp *types.GetTopicRegistrationsResponse, err error) {
	l.logHelper.Info("GetListTopicRegistationLogic", req)
	var data []types.TopicRegistration
	mapFaculties := make(map[int32]string, 0)
	mapUserInfo := make(map[int32]db.UserInfo, 0)
	mapUser := make(map[int32]int32, 0)

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
	users, err := l.svcCtx.Store.ListUsers(l.ctx)
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
		mapUser[user.ID] = user.TypeAccount
	}

	lectures, err := l.svcCtx.Store.ListUserInfosByName(l.ctx, fmt.Sprintf("%%%s%%", req.Lecture))
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetTopicRegistrationsResponse{
				Result: types.Result{
					Code:    constant.USER_NOT_FOUND_CODE,
					Message: constant.USER_NOT_FOUND_MESSAGE,
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
	for _, lecture := range lectures {
		if mapUser[lecture.UserID] == constant.LectureType {
			mapUserInfo[lecture.UserID] = lecture
		}
	}
	for _, tmp := range list {
		if req.Status != 0 && req.Status != tmp.Status {
			continue
		}
		if req.FacultyID != 0 && req.FacultyID != tmp.FacultyID {
			continue
		}
		if _, ok := mapUserInfo[tmp.LectureID]; !ok {
			continue
		}
		data = append(data, types.TopicRegistration{
			ID:      tmp.ID,
			Name:    tmp.Name,
			Lecture: mapUserInfo[tmp.LectureID].Name,
			Email:   mapUserInfo[tmp.LectureID].Email,
			Phone:   mapUserInfo[tmp.LectureID].Phone,
			Faculty: mapFaculties[tmp.FacultyID],
			Status:  constant.MapStatusTopicRegistration[tmp.Status],
		})
	}
	if req.Limit != 0 {
		data = utils.SliceArray(data, req.Limit, req.Offset)
	}

	return &types.GetTopicRegistrationsResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Total:              len(data),
		TopicRegistrations: data,
	}, nil
}

func (l *Logic) CreateTopicRegistationLogic(lectureID int32, req *types.CreateTopicRegistrationRequest) (resp *types.CreateTopicRegistrationResponse, err error) {
	l.logHelper.Info("CreateTopicRegistationLogic ", req)

	lecture, err := l.svcCtx.Store.GetUserInfo(l.ctx, lectureID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.CreateTopicRegistrationResponse{
				Result: types.Result{
					Code:    constant.USER_NOT_FOUND_CODE,
					Message: constant.USER_NOT_FOUND_MESSAGE,
				},
			}, nil
		}
	}

	_, err = l.svcCtx.Store.CreateTopicRegistration(l.ctx, db.CreateTopicRegistrationParams{
		ID:        utils.RandomID(),
		Name:      req.Name,
		LectureID: lectureID,
		FacultyID: lecture.FacultyID,
		Status:    1,
		CreatedAt: time.Now(),
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

	_, err = l.svcCtx.Store.GetTopicRegistration(l.ctx, id)
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
		Name:      req.Name,
		LectureID: req.FacultyID,
		FacultyID: req.FacultyID,
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
			Email:   lecture.Email,
			Phone:   lecture.Phone,
			Faculty: faculty.Name,
			Status:  constant.MapStatusTopicRegistration[topicRegis.Status],
		},
	}, nil
}

func (l *Logic) AcceptTopicRegistationLogic(req *types.AccceptTopicRegistrationRequest) (resp *types.AcceptTopicRegistrationResponse, err error) {
	l.logHelper.Info("AcceptTopicRegistationLogic", req)

	var listTopicRegistrationId []int32

	err = json.Unmarshal([]byte(req.ListTopicRegistrationId), &listTopicRegistrationId)
	if err != nil {
		l.logHelper.Error(err)
		return &types.AcceptTopicRegistrationResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	for _, id := range listTopicRegistrationId {
		topicRegis, err := l.svcCtx.Store.GetTopicRegistration(l.ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return &types.AcceptTopicRegistrationResponse{
					Result: types.Result{
						Code:    constant.TOPIC_REGIS_NOT_FOUND_CODE,
						Message: constant.TOPIC_REGIS_NOT_FOUND_MESSAGE,
					},
				}, nil
			}
			l.logHelper.Error(err)
			return &types.AcceptTopicRegistrationResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}
		err = l.svcCtx.Store.AcceptTopicRegistration(l.ctx, db.AcceptTopicRegistrationParams{
			ID:     topicRegis.ID,
			Status: req.Status,
		})
		if err != nil {
			l.logHelper.Error(err)
			return &types.AcceptTopicRegistrationResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}
	}

	return &types.AcceptTopicRegistrationResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}
