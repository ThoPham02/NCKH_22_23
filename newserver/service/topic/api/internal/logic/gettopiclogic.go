package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	userModel "github.com/ThoPham02/research_management/service/account/model"
	subcommitteeModel "github.com/ThoPham02/research_management/service/result/model"
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
	var topic types.TopicDetail
	var subcommittee subcommitteeModel.SubcommitteeTbl
	var listStudent []types.StudentInfo
	var listStudentModel []userModel.UserTbl

	topicModel, err = l.svcCtx.TopicModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.GetTopicRes{
				Result: types.Result{
					Code:    common.SUCCESS_CODE,
					Message: common.SUCCESS_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, topicModel.LectureId)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	department, err := l.svcCtx.DepartmentModel.FindOne(l.ctx, topicModel.DepartmentId)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	faculty, err := l.svcCtx.FacultyModel.FindOne(l.ctx, department.FacultyId)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if topicModel.SubcommitteeId.Int64 != 0 {
		subcommitteeModel, err := l.svcCtx.SubcommitteeModel.FindOne(l.ctx, topicModel.SubcommitteeId.Int64)
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
			subcommittee = *subcommitteeModel
		}
	}

	if topicModel.GroupStudentsId.Int64 != 0 {
		listStudentModel, err = l.svcCtx.StudentGroupModel.FindStudentByGroupID(l.ctx, topicModel.GroupStudentsId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetTopicRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		for _, tmp := range listStudentModel {
			listStudent = append(listStudent, types.StudentInfo{
				ID:    tmp.Id,
				Name:  tmp.Name,
				Email: tmp.Email.String,
				Phone: tmp.Phone.String,
			})
		}
	}

	event, err := l.svcCtx.EventModel.FindOne(l.ctx, topicModel.EventId)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetTopicRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	topic = types.TopicDetail{
		ID:           topicModel.Id,
		Name:         topicModel.Name,
		LectureInfo:  types.LectureInfo{ID: user.Id, Name: user.Name, Email: user.Email.String, Phone: user.Phone.String, Degree: common.MapDegree[user.Degree]},
		Department:   department.Name,
		Faculty:      faculty.Name,
		Status:       topicModel.Status,
		Subcommittee: subcommittee.Name.String,
		ListStudent:  listStudent,
		TimeStart:    topicModel.TimeStart.Int64,
		TimeEnd:      topicModel.TimeEnd.Int64,
		Event:        event.Name,
		CashSupport:  topicModel.CashSupport.Int64,
	}

	return &types.GetTopicRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		TopicDetail: topic,
	}, nil
}
