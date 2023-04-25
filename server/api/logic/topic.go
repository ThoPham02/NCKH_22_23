package logic

import (
	"database/sql"
	"encoding/json"
	"github/ThoPham02/research_management/api/constant"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
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

func (l *Logic) GetTopicLogic(req *types.GetTopicRequest) (resp *types.GetTopicResponse, err error) {

	return
}

func (l *Logic) AcceptTopicLogic(req *types.AcceptTopicRequest) (resp *types.AcceptTopicResponse, err error) {
	return
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

	timeStart, err := utils.ConvertStringToTime(req.TimeStart)
	if err != nil {
		l.logHelper.Error(err)
		return &types.CreateTopicResponse{
			Result: types.Result{
				Code:    constant.DB_ERR_CODE,
				Message: constant.DB_ERR_MESSAGE,
			},
		}, nil
	}
	timeEnd, err := utils.ConvertStringToTime(req.TimeEnd)
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
		TimeStart:    *timeStart,
		TimeEnd:      *timeEnd,
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

func (l *Logic) UpdateTopicLogic(id int32, req *types.UpdateTopicRequest) (resp *types.UpdateTopicResponse, err error) {
	return
}
