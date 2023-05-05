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
	l.logHelper.Info(listStudents)
	if err != nil {
		l.logHelper.Error(err)
		return &types.GetTopicByIdResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	for _, tmp := range listStudents {
		name, err := l.svcCtx.Store.GetUserNameByID(l.ctx, tmp)
		if err != nil {
			l.logHelper.Error(err)
			return &types.GetTopicByIdResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}
		listName = append(listName, name)
	}

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
			ListStudents: listName,
			TimeStart:    data.TimeStart.Format(constant.Layout),
			TimeEnd:      data.TimeEnd.Format(constant.Layout),
		},
	}, nil
}

func (l *Logic) ListTopicLogic(req *types.GetTopicRequest) (resp *types.GetTopicResponse, err error) {
	l.logHelper.Info("Get Topic Logic", req)
	topics, err := l.svcCtx.Store.ListTopicsFilter(l.ctx, fmt.Sprintf("%%%s%%", req.Search))
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetTopicResponse{
				Result: types.Result{
					Code:    constant.SUCCESS_CODE,
					Message: constant.SUCCESS_MESSAGE,
				},
			}, nil
		}
	}

	var list []types.Topic

	for _, topic := range topics {
		// if req.FacultyID != 0 && req.FacultyID != topic.FacultyID {
		// 	continue
		// }
		// if req.LectureID != 0 && req.LectureID != topic.LectureID {
		// 	continue
		// }

		// var listStudent []string
		// studentID, err := l.svcCtx.Store.ListStudentByTopicId(l.ctx, topic.ID)
		// if err != nil {
		// 	l.logHelper.Error(err)
		// 	return &types.GetTopicResponse{
		// 		Result: types.Result{
		// 			Code:    constant.DB_ERR_CODE,
		// 			Message: constant.DB_ERR_MESSAGE,
		// 		},
		// 	}, nil
		// }
		// for _, id := range studentID {
		// 	student, err := l.svcCtx.Store.GetUser(l.ctx, id)
		// 	if err != nil {
		// 		l.logHelper.Error(err)
		// 		return &types.GetTopicResponse{
		// 			Result: types.Result{
		// 				Code:    constant.DB_ERR_CODE,
		// 				Message: constant.DB_ERR_MESSAGE,
		// 			},
		// 		}, nil
		// 	}

		// 	listStudent = append(listStudent, student.Name)
		// }
		// if req.StudentID != 0 {
		// 	continue
		// }

		lecture, err := l.svcCtx.Store.GetUserNameByID(l.ctx, topic.LectureID)
		if err != nil {
			l.logHelper.Error(err)
			return &types.GetTopicResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}

		faculty, err := l.svcCtx.Store.GetFaculty(l.ctx, topic.FacultyID)
		if err != nil {
			l.logHelper.Error(err)
			return &types.GetTopicResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}

		list = append(list, types.Topic{
			ID:           topic.ID,
			Name:         topic.Name,
			Lecture:      lecture,
			Faculty:      faculty.Name,
			Status:       constant.MapStatusTopic[topic.Status],
			ResultUrl:    topic.ResultUrl.String,
			ListStudents: []string{},
			TimeStart:    topic.TimeStart.Format(constant.Layout),
			TimeEnd:      topic.TimeEnd.Format(constant.Layout),
		})
	}

	return &types.GetTopicResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
		Total:      len(list),
		ListTopics: list,
	}, nil
}

func (l *Logic) AcceptTopicLogic(req *types.AcceptTopicRequest) (resp *types.AcceptTopicResponse, err error) {
	l.logHelper.Info("AcceptTopicLogic", req)

	var listTopicId []int32

	err = json.Unmarshal([]byte(req.ListTopicID), &listTopicId)
	if err != nil {
		l.logHelper.Error(err)
		return &types.AcceptTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	timeStart, err := utils.ConvertStringToTime(req.TimeStart)
	if err != nil {
		l.logHelper.Error(err)
		return &types.AcceptTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	timeEnd, err := utils.ConvertStringToTime(req.TimeEnd)
	if err != nil {
		l.logHelper.Error(err)
		return &types.AcceptTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	for _, id := range listTopicId {
		topic, err := l.svcCtx.Store.GetTopic(l.ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return &types.AcceptTopicResponse{
					Result: types.Result{
						Code:    constant.TOPIC_NOT_FOUND_CODE,
						Message: constant.TOPIC_NOT_FOUND_MESS,
					},
				}, nil
			}
		}

		err = l.svcCtx.Store.AcceptTopic(l.ctx, db.AcceptTopicParams{
			ID:           id,
			Name:         topic.Name,
			LectureID:    topic.LectureID,
			FacultyID:    topic.FacultyID,
			Status:       2,
			ResultUrl:    sql.NullString{Valid: false},
			ConferenceID: topic.ConferenceID,
			GroupID:      sql.NullInt32{Valid: false},
			TimeStart:    *timeStart,
			TimeEnd:      *timeEnd,
		})
		if err != nil {
			l.logHelper.Error(err)
			return &types.AcceptTopicResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}
	}

	return &types.AcceptTopicResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}

func (l *Logic) CreateTopicLogic(req *types.CreateTopicRequest) (resp *types.CreateTopicResponse, err error) {
	l.logHelper.Info("CreateTopicLogic", req)

	var listStudent []int32

	err = json.Unmarshal([]byte(req.ListStudent), &listStudent)
	if err != nil {
		l.logHelper.Error(err)
		return &types.CreateTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	topicRegis, err := l.svcCtx.Store.GetTopicRegistration(l.ctx, req.TopicRegisID)
	if err != nil {
		l.logHelper.Error(err)
		return &types.CreateTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	topic, err := l.svcCtx.Store.CreateTopic(l.ctx, db.CreateTopicParams{
		ID:           utils.RandomID(),
		Name:         topicRegis.Name,
		LectureID:    topicRegis.LectureID,
		FacultyID:    topicRegis.FacultyID,
		Status:       1,
		ResultUrl:    sql.NullString{Valid: false},
		ConferenceID: req.ConferenceID,
		GroupID:      sql.NullInt32{Valid: false},
		TimeStart:    time.Time{},
		TimeEnd:      time.Time{},
	})
	if err != nil {
		l.logHelper.Error(err)
		return &types.CreateTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	for _, studentId := range listStudent {
		_, err := l.svcCtx.Store.CreateStudentTopic(l.ctx, db.CreateStudentTopicParams{
			ID:        utils.RandomID(),
			TopicID:   topic.ID,
			StudentID: studentId,
		})
		if err != nil {
			return &types.CreateTopicResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}
	}

	return &types.CreateTopicResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}

func (l *Logic) UpdateGroupTopic(req *types.UpdateTopicRequest) (resp *types.UpdateTopicResponse, err error) {
	l.logHelper.Info("UpdateTopicLogic", req)

	var listTopicID []int32

	err = json.Unmarshal([]byte(req.ListTopicID), &listTopicID)
	if err != nil {
		l.logHelper.Error(err)
		return &types.UpdateTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}

	for _, id := range listTopicID {
		_, err = l.svcCtx.Store.GetTopic(l.ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return &types.UpdateTopicResponse{
					Result: types.Result{
						Code:    constant.TOPIC_NOT_FOUND_CODE,
						Message: constant.TOPIC_NOT_FOUND_MESS,
					},
				}, nil
			}
		}

		err = l.svcCtx.Store.UpdateGroupTopic(l.ctx, db.UpdateGroupTopicParams{
			ID:      id,
			GroupID: utils.GetInt32(req.GroupID),
		})
		if err != nil {
			l.logHelper.Error(err)
			return &types.UpdateTopicResponse{
				Result: types.Result{
					Code:    constant.DB_ERR_CODE,
					Message: constant.DB_ERR_MESSAGE,
				},
			}, nil
		}
	}

	return &types.UpdateTopicResponse{
		Result: types.Result{
			Code:    constant.SUCCESS_CODE,
			Message: constant.SUCCESS_MESSAGE,
		},
	}, nil
}
