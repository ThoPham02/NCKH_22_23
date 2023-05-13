package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TopicTblModel = (*customTopicTblModel)(nil)

type TopicConditions struct {
	Search         string `json:"search"`
	DepartmentID   int64  `json:"departmentID"`
	FacultyID      int64  `json:"facultyID"`
	Status         int64  `json:"status"`
	LectureID      int64  `json:"lectureID"`
	EventID        int64  `json:"eventID"`
	SubcommitteeID int64  `json:"subcommitteeID"`
	TimeStart      int64  `json:"timeStart"`
	TimeEnd        int64  `json:"timeEnd"`
	Limit          int64  `json:"limit"`
	Offset         int64  `json:"offset"`
}

type (
	// TopicTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicTblModel.
	TopicTblModel interface {
		topicTblModel
		FindTopics(ctx context.Context, condition TopicConditions) ([]TopicTbl, error)
		CountTopics(ctx context.Context, condition TopicConditions) (int64, error)
		UpdateGroup(ctx context.Context, topicID int64, groupID int64) error
		UpdateStatus(ctx context.Context, topicID int64, status int64) error
		UpdateSubcommittee(ctx context.Context, topicID []int64, subcommitteeID int64) error
	}

	customTopicTblModel struct {
		*defaultTopicTblModel
	}
)

// NewTopicTblModel returns a model for the database table.
func NewTopicTblModel(conn sqlx.SqlConn) TopicTblModel {
	return &customTopicTblModel{
		defaultTopicTblModel: newTopicTblModel(conn),
	}
}

func (m *customTopicTblModel) FindTopics(ctx context.Context, condition TopicConditions) ([]TopicTbl, error) {
	var data []TopicTbl
	var values = []interface{}{fmt.Sprintf("%%%s%%", condition.Search)}
	query := fmt.Sprintf("select %s from %s where name ilike $1", topicTblRows, m.table)
	if condition.DepartmentID != 0 {
		values = append(values, condition.DepartmentID)
		query += fmt.Sprintf(" and department_id = $%d", len(values))
	}
	if condition.FacultyID != 0 {
		values = append(values, condition.FacultyID)
		query += fmt.Sprintf(" and department_id in (select id from department_tbl where faculty_id = $%d)", len(values))
	}
	if condition.LectureID != 0 {
		values = append(values, condition.LectureID)
		query += fmt.Sprintf(" and lecture_id = $%d", len(values))
	}
	if condition.Status != 0 {
		values = append(values, condition.Status)
		query += fmt.Sprintf(" and status = $%d", len(values))
	}
	if condition.EventID != 0 {
		values = append(values, condition.EventID)
		query += fmt.Sprintf(" and event_id = $%d", len(values))
	}
	if condition.SubcommitteeID != 0 {
		values = append(values, condition.SubcommitteeID)
		query += fmt.Sprintf(" and subcommittee_id = $%d", len(values))
	}
	if condition.TimeStart > 0 {
		values = append(values, condition.TimeStart)
		query += fmt.Sprintf(" and time_start >= $%d", len(values))
	}
	if condition.TimeEnd > 0 {
		values = append(values, condition.TimeEnd)
		query += fmt.Sprintf(" and time_end <= $%d", len(values))
	}
	if condition.Limit > 0 {
		values = append(values, condition.Limit, condition.Offset)
		query += fmt.Sprintf(" limit $%d offset $%d", len(values)-1, len(values))
	}

	err := m.conn.QueryRowsCtx(ctx, &data, query, values...)
	switch err {
	case nil:
		return data, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTopicTblModel) CountTopics(ctx context.Context, condition TopicConditions) (int64, error) {
	var data int64
	var values = []interface{}{fmt.Sprintf("%%%s%%", condition.Search)}
	query := fmt.Sprintf("select count(*) from %s where name ilike $1", m.table)
	if condition.DepartmentID != 0 {
		values = append(values, condition.DepartmentID)
		query += fmt.Sprintf(" and department_id = $%d", len(values))
	}
	if condition.FacultyID != 0 {
		values = append(values, condition.FacultyID)
		query += fmt.Sprintf(" and department_id in (select id from department_tbl where faculty_id = $%d)", len(values))
	}
	if condition.LectureID != 0 {
		values = append(values, condition.LectureID)
		query += fmt.Sprintf(" and lecture_id = $%d", len(values))
	}
	if condition.Status != 0 {
		values = append(values, condition.Status)
		query += fmt.Sprintf(" and status = $%d", len(values))
	}
	if condition.EventID != 0 {
		values = append(values, condition.EventID)
		query += fmt.Sprintf(" and event_id = $%d", len(values))
	}
	if condition.SubcommitteeID != 0 {
		values = append(values, condition.SubcommitteeID)
		query += fmt.Sprintf(" and subcommittee_id = $%d", len(values))
	}
	if condition.TimeStart > 0 {
		values = append(values, condition.TimeStart)
		query += fmt.Sprintf(" and time_start >= $%d", len(values))
	}
	if condition.TimeEnd > 0 {
		values = append(values, condition.TimeEnd)
		query += fmt.Sprintf(" and time_end <= $%d", len(values))
	}

	err := m.conn.QueryRowCtx(ctx, &data, query, values...)
	switch err {
	case nil:
		return data, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
func (m *customTopicTblModel) UpdateStatus(ctx context.Context, topicID int64, status int64) error {
	logx.Info(topicTblRowsWithPlaceHolder)
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, "status = $2")
	_, err := m.conn.ExecCtx(ctx, query, topicID, status)
	return err
}

func (m *customTopicTblModel) UpdateGroup(ctx context.Context, topicID int64, groupID int64) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, "group_students_id = $2")
	_, err := m.conn.ExecCtx(ctx, query, topicID, groupID)
	return err
}
func (m *customTopicTblModel) UpdateSubcommittee(ctx context.Context, topicID []int64, subcommitteeID int64) error {
	var input string = ""
	var values = []interface{}{subcommitteeID}
	for index, id := range topicID {
		input += fmt.Sprintf("$%d, ", index+2)
		values = append(values, id)
	}
	input = input[:len(input)-2]
	query := fmt.Sprintf("update %s set %s where id in (%s)", m.table, "subcommittee_id = $1", input)
	_, err := m.conn.ExecCtx(ctx, query, values...)
	return err
}
