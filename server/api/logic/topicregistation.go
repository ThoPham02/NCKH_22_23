package logic

import (
	"database/sql"
	"errors"
	"fmt"
	db "github/ThoPham02/research_management/api/db/sqlc"
	"github/ThoPham02/research_management/api/types"
	"github/ThoPham02/research_management/api/utils"
	"time"
)

func (l *Logic) GetListTopicRegistationLogic(req *types.GetListTopicRegistationRequest) (*types.GetListTopicRegistationResponse, error) {
	l.logHelper.Info("GetListTopicRegistationLogic", req)

	var err error
	var list []types.TopicRegistation

	if req == nil {
		err = errors.New("invalid request")
		l.logHelper.Error(err)
		return nil, err
	}

	data, err := l.svcCtx.Store.GetListTopicRegistation(l.ctx, fmt.Sprintf("%%%s%%", req.Search))
	if err != nil {
		if err == sql.ErrNoRows {
			return &types.GetListTopicRegistationResponse{
				Total: 0,
			}, nil
		}
		l.logHelper.Error(err)
		return nil, err
	}

	if req.FaculityID != 0 {
		data = utils.FilterByField(data, "FaculityID", req.FaculityID).([]db.TopicRegistration)
	}
	if req.LectureID != 0 {
		data = utils.FilterByField(data, "LectureID", req.FaculityID).([]db.TopicRegistration)
	}

	total := len(data)

	if req.Limit != 0 {
		data = utils.SliceArray(data, req.Limit, req.Offset-1)
	}

	for _, item := range data {
		list = append(list, types.TopicRegistation{
			ID:             item.ID,
			Name:           item.Name,
			Description:    item.Description,
			DescriptionURL: item.DescriptionUrl.String,
			LectureID:      item.LectureID,
			FaculityID:     item.FaculityID.Int64,
			CreatedAt:      item.CreatedAt.Time.String(),
		})
	}

	return &types.GetListTopicRegistationResponse{
		Total:                int64(total),
		ListTopicRegistation: list,
	}, nil
}

func (l *Logic) CreateTopicRegistation(req *types.CreateTopicRegistationRequest) (*types.CreateTopicRegistationResponse, error) {
	l.logHelper.Info("CreateTopicRegistation ", req)

	var err error

	if req == nil {
		err = errors.New("invalid request")
		l.logHelper.Error(err)
		return nil, err
	}

	data, err := l.svcCtx.Store.CreateTopicRegistation(l.ctx, db.CreateTopicRegistationParams{
		Name:        req.Name,
		Description: req.Description,
		DescriptionUrl: sql.NullString{
			Valid:  req.DescriptionUrl != "",
			String: req.DescriptionUrl,
		},
		LectureID: req.LectureID,
		FaculityID: sql.NullInt64{
			Valid: req.FaculityID != 0,
			Int64: req.FaculityID,
		},
		CreatedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	})
	if err != nil {
		l.logHelper.Error(err)
		return nil, err
	}

	return &types.CreateTopicRegistationResponse{
		ID:             data.ID,
		Name:           data.Name,
		Description:    data.Description,
		DescriptionUrl: data.DescriptionUrl.String,
		LectureID:      data.LectureID,
		FaculityID:     data.FaculityID.Int64,
	}, nil
}

func (l *Logic) UpdateTopicRegistationLogic(req *types.UpdateTopicRegistationRequest) (*types.UpdateTopicRegistationResponse, error) {
	l.logHelper.Info("UpdateTopicRegistation", req)

	var err error

	if req == nil {
		err = errors.New("invalid request")
		l.logHelper.Error(err)
		return nil, err
	}

	data, err := l.svcCtx.Store.GetTopicRegistationById(l.ctx, req.ID)
	if err != nil {
		l.logHelper.Error(err)
		return nil, err
	}

	res, err := l.svcCtx.Store.UpdateTopicRegistation(l.ctx, db.UpdateTopicRegistationParams{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		DescriptionUrl: sql.NullString{
			Valid:  req.DescriptionUrl != "",
			String: req.DescriptionUrl,
		},
		LectureID:  req.LectureID,
		FaculityID: data.FaculityID,
	})
	if err != nil {
		l.logHelper.Error(err)
		return nil, err
	}

	return &types.UpdateTopicRegistationResponse{
		ID:             res.ID,
		Name:           res.Name,
		Description:    res.Description,
		DescriptionUrl: res.DescriptionUrl.String,
		LectureID:      res.LectureID,
		FaculityID:     res.FaculityID.Int64,
	}, nil
}

func (l *Logic) GetTopicRegistationByIdLogic(req *types.GetTopicRegistationByIDRequest) (*types.GetTopicRegistationByIdResponse, error) {
	l.logHelper.Info("GetTopicRegistation", req)
	data, err := l.svcCtx.Store.GetTopicRegistationById(l.ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		l.logHelper.Error(err)
		return nil, err
	}
	return &types.GetTopicRegistationByIdResponse{
		ID:             data.ID,
		Name:           data.Name,
		Description:    data.Description,
		DescriptionUrl: data.DescriptionUrl.String,
		LectureID:      data.LectureID,
		FaculityID:     data.FaculityID.Int64,
	}, err
}
