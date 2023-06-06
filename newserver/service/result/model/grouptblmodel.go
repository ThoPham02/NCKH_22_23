package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupTblModel = (*customGroupTblModel)(nil)

type (
	// GroupTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupTblModel.
	GroupTblModel interface {
		groupTblModel
		InsertMutil(ctx context.Context, groupsModel []GroupTbl) error
		FindMulti(ctx context.Context) ([]GroupTbl, error)
		FindMultiBySubcommittee(ctx context.Context, listSub []int64) ([]GroupTbl, error)
	}

	customGroupTblModel struct {
		*defaultGroupTblModel
	}
)

// NewGroupTblModel returns a model for the database table.
func NewGroupTblModel(conn sqlx.SqlConn) GroupTblModel {
	return &customGroupTblModel{
		defaultGroupTblModel: newGroupTblModel(conn),
	}
}

func (m *customGroupTblModel) InsertMutil(ctx context.Context, data []GroupTbl) error {
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, groupTblRowsExpectAutoSet)
	var values = []interface{}{}
	var tmp int
	for _, group := range data {
		tmp = len(values)
		query += fmt.Sprintf("($%d, $%d, $%d, $%d), ", tmp+1, tmp+2, tmp+3, tmp+4)
		values = append(values, group.Id, group.SubcommitteeId, group.LectureId, group.Role)
	}
	query = query[0 : len(query)-2]
	_, err := m.conn.ExecCtx(ctx, query, values...)
	return err
}

func (m *customGroupTblModel) FindMulti(ctx context.Context) ([]GroupTbl, error) {
	query := fmt.Sprintf("select %s from %s ", groupTblRows, m.table)
	var resp []GroupTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customGroupTblModel) FindMultiBySubcommittee(ctx context.Context, listSub []int64) ([]GroupTbl, error) {
	if len(listSub) == 0 {
		return nil, nil
	}
	var listString = ""
	for _, tmp := range listSub {
		listString += fmt.Sprintf("%d,", tmp)
	}
	listString = listString[0 : len(listString)-1]
	query := fmt.Sprintf("select %s from %s where subcommittee_id in (%s)", groupTblRows, m.table, listString)
	var resp []GroupTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
