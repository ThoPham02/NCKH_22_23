package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	userModel "github.com/ThoPham02/research_management/service/account/model"
	progressModel "github.com/ThoPham02/research_management/service/progress/model"
	resultModel "github.com/ThoPham02/research_management/service/result/model"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/ThoPham02/research_management/service/topic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicLogic {
	return &GetTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicLogic) GetTopic(req *types.GetTopicReq) (resp *types.GetTopicRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetTopic", req)

	var topicModel *model.TopicTbl
	var subcommitteeModel *resultModel.SubcommitteeTbl
	var groupStudents []userModel.UserTbl
	var topicMarksModel []resultModel.TopicMarkTbl
	var reportsModel []progressModel.TopicReportTbl

	var lectureInfo *types.LectureInfo
	var topic types.Topic
	var event types.Event
	var subcommittee types.Subcommittee
	var reports []types.Report
	var marks []types.Mark
	var listStudent []types.Student

	topicModel, err = l.svcCtx.TopicModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if topicModel == nil {
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.TOPIC_NOT_EXIST_CODE,
				Message: common.TOPIC_NOT_EXIST_MESS,
			},
		}, nil
	}

	lectureInfo, err = l.getLectureInfoById(topicModel.LectureId)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	topic = types.Topic{
		ID:              topicModel.Id,
		Name:            topicModel.Name,
		LectureInfo:     *lectureInfo,
		DepartmentID:    topicModel.DepartmentId,
		Status:          topicModel.Status,
		EventId:         topicModel.EventId,
		SubcommitteeID:  topicModel.SubcommitteeId.Int64,
		TimeStart:       topicModel.TimeStart.Int64,
		TimeEnd:         topicModel.TimeEnd.Int64,
		GroupStudentId:  topicModel.GroupStudentsId.Int64,
		EstimateStudent: topicModel.EstimateStudent,
	}

	if topicModel.SubcommitteeId.Valid {
		subcommitteeModel, err = l.svcCtx.SubcommitteeModel.FindOne(l.ctx, topicModel.SubcommitteeId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetTopicRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		if subcommitteeModel != nil {
			subcommittee = types.Subcommittee{
				ID:        subcommitteeModel.Id,
				Name:      subcommitteeModel.Name,
				FacultyID: subcommitteeModel.FacultId.Int64,
				EventID:   subcommitteeModel.EventId,
				Level:     subcommitteeModel.Level,
			}
		}
	}

	if topicModel.GroupStudentsId.Valid {
		groupStudents, err = l.svcCtx.StudentGroupModel.FindStudentByGroupID(l.ctx, topicModel.GroupStudentsId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetTopicRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		if len(groupStudents) > 0 {
			for _, tmp := range groupStudents {
				listStudent = append(listStudent, types.Student{
					StudentID: tmp.Id,
					Name:      tmp.Name,
					EventID:   topicModel.EventId,
					GroupID:   topicModel.GroupStudentsId.Int64,
				})
			}
		}
	}
	topicMarksModel, err = l.svcCtx.TopicMarkModel.FindByTopicID(l.ctx, topicModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, topicMarkModel := range topicMarksModel {
		lectureInfo, err = l.getLectureInfoById(topicMarkModel.LectureId)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetTopicRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		marks = append(marks, types.Mark{
			ID:          topicMarkModel.Id,
			TopicID:     topicMarkModel.TopicId,
			LectureInfo: *lectureInfo,
			Point:       topicMarkModel.Point,
			Comment:     topicMarkModel.Comment.String,
			Url:         topicMarkModel.Url.String,
			Level:       topicMarkModel.Level,
			CreatedAt:   topicMarkModel.CreatedAt,
		})
	}

	reportsModel, err = l.svcCtx.TopicReportModel.FindByTopicID(l.ctx, topicModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, tmp := range reportsModel {
		reports = append(reports, types.Report{
			ID:          tmp.Id,
			TopicID:     tmp.TopicId,
			StageID:     tmp.StageId,
			Description: tmp.Description.String,
			CreatedAt:   tmp.CreatedAt,
			CreatedBy:   tmp.CreatedBy,
		})
	}

	eventModel, err := l.svcCtx.EventModel.FindOne(l.ctx, topicModel.EventId)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	event = types.Event{
		ID:         eventModel.Id,
		Name:       eventModel.Name,
		SchoolYear: eventModel.SchoolYear.String,
		IsCurrent:  eventModel.IsCurrent.Int64,
	}

	return &types.GetTopicRes{
		Result:       types.Result{Code: common.SUCCESS_CODE, Message: common.SUCCESS_MESS},
		Topic:        topic,
		Event:        event,
		Subcommittee: subcommittee,
		Reports:      reports,
		Marks:        marks,
		ListStudent:  listStudent,
	}, nil
}

func (l *GetTopicLogic) getLectureInfoById(lectureID int64) (*types.LectureInfo, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, lectureID)
	if err != nil {
		return nil, err
	}
	return &types.LectureInfo{
		ID:     lectureID,
		Name:   user.Name,
		Email:  user.Email.String,
		Phone:  user.Phone.String,
		Degree: common.MapDegree[user.Degree],
	}, nil
}
