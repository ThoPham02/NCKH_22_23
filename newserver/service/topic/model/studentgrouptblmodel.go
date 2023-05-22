package model

import (
	"context"
	"fmt"
	"strings"

	userModel "github.com/ThoPham02/research_management/service/account/model"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudentGroupTblModel = (*customStudentGroupTblModel)(nil)

var (
	userTblFieldNames = builder.RawFieldNames(&userModel.UserTbl{}, true)
	userTblRows       = strings.Join(userTblFieldNames, ",")
)

type (
	// StudentGroupTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentGroupTblModel.
	StudentGroupTblModel interface {
		studentGroupTblModel
		InsertMutil(ctx context.Context, data []StudentGroupTbl) error
		FindStudentByGroupID(ctx context.Context, groupId int64) ([]userModel.UserTbl, error)
		CheckStudentValid(ctx context.Context, studentId int64, eventId int64) (bool, error)
		DeleteByStudentID(ctx context.Context, studentId int64, eventId int64) error
	}

	customStudentGroupTblModel struct {
		*defaultStudentGroupTblModel
	}
)

// NewStudentGroupTblModel returns a model for the database table.
func NewStudentGroupTblModel(conn sqlx.SqlConn) StudentGroupTblModel {
	return &customStudentGroupTblModel{
		defaultStudentGroupTblModel: newStudentGroupTblModel(conn),
	}
}

func (m *customStudentGroupTblModel) InsertMutil(ctx context.Context, data []StudentGroupTbl) error {
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, studentGroupTblRows)
	var values = []interface{}{}
	var tmp int
	for _, group := range data {
		tmp = len(values)
		query += fmt.Sprintf("($%d, $%d, $%d, $%d), ", tmp+1, tmp+2, tmp+3, tmp+4)
		values = append(values, group.Id, group.EventId, group.StudentId, group.GroupId)
	}
	query = query[0 : len(query)-2]
	_, err := m.conn.ExecCtx(ctx, query, values...)
	return err
}

func (m *customStudentGroupTblModel) FindStudentByGroupID(ctx context.Context, groupId int64) ([]userModel.UserTbl, error) {
	query := fmt.Sprintf("select %s from %s where id in (select student_id from %s where group_id = $1)", userTblRows, "user_tbl", m.table)
	var resp []userModel.UserTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, groupId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customStudentGroupTblModel) CheckStudentValid(ctx context.Context, studentId int64, eventId int64) (bool, error) {
	query := fmt.Sprintf("select %s from %s where student_id = $1 and event_id = $2 limit 1", studentGroupTblRows, m.table)
	var resp StudentGroupTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, studentId, eventId)
	switch err {
	case nil:
		return false, nil
	case sqlc.ErrNotFound:
		return true, nil
	default:
		return false, err
	}
}

func (m *customStudentGroupTblModel) DeleteByStudentID(ctx context.Context, studentId int64, eventId int64) error {
	query := fmt.Sprintf("delete from %s where student_id = $1 and event_id = $2", m.table)
	_, err := m.conn.ExecCtx(ctx, query, studentId, eventId)
	return err
}
