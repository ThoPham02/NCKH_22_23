// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: topic.sql

package db

import (
	"context"
	"database/sql"
)

const createTopic = `-- name: CreateTopic :one
INSERT INTO
    topics (
        name,
        description,
        department_id,
        time_start,
        time_end,
        deadline,
        status_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, description, department_id, time_start, time_end, deadline, status_id, created_at, updated_at, deleted_at
`

type CreateTopicParams struct {
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	DepartmentID int64          `json:"department_id"`
	TimeStart    sql.NullTime   `json:"time_start"`
	TimeEnd      sql.NullTime   `json:"time_end"`
	Deadline     sql.NullTime   `json:"deadline"`
	StatusID     int64          `json:"status_id"`
}

func (q *Queries) CreateTopic(ctx context.Context, arg CreateTopicParams) (Topic, error) {
	row := q.db.QueryRowContext(ctx, createTopic,
		arg.Name,
		arg.Description,
		arg.DepartmentID,
		arg.TimeStart,
		arg.TimeEnd,
		arg.Deadline,
		arg.StatusID,
	)
	var i Topic
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DepartmentID,
		&i.TimeStart,
		&i.TimeEnd,
		&i.Deadline,
		&i.StatusID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteTopic = `-- name: DeleteTopic :exec
Update
    topics
SET
    deleted_at = NOW()
WHERE
    id = $1
`

func (q *Queries) DeleteTopic(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTopic, id)
	return err
}

const getTopic = `-- name: GetTopic :one
SELECT
    id, name, description, department_id, time_start, time_end, deadline, status_id, created_at, updated_at, deleted_at
FROM
    topics
WHERE
    deleted_at IS NULL
    AND id = $1
`

func (q *Queries) GetTopic(ctx context.Context, id int64) (Topic, error) {
	row := q.db.QueryRowContext(ctx, getTopic, id)
	var i Topic
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DepartmentID,
		&i.TimeStart,
		&i.TimeEnd,
		&i.Deadline,
		&i.StatusID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listTopics = `-- name: ListTopics :many
SELECT
    id, name, description, department_id, time_start, time_end, deadline, status_id, created_at, updated_at, deleted_at
FROM
    topics
WHERE
    deleted_at IS NULL
ORDER BY
    id
LIMIT $1 OFFSET $2
`

type ListTopicsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTopics(ctx context.Context, arg ListTopicsParams) ([]Topic, error) {
	rows, err := q.db.QueryContext(ctx, listTopics, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Topic{}
	for rows.Next() {
		var i Topic
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.DepartmentID,
			&i.TimeStart,
			&i.TimeEnd,
			&i.Deadline,
			&i.StatusID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTopic = `-- name: UpdateTopic :one
UPDATE
    topics
SET
    name = $2,
    description = $3,
    department_id = $4,
    time_start = $5,
    time_end = $6,
    deadline = $7,
    status_id = $8
WHERE
    id = $1 RETURNING id, name, description, department_id, time_start, time_end, deadline, status_id, created_at, updated_at, deleted_at
`

type UpdateTopicParams struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	DepartmentID int64          `json:"department_id"`
	TimeStart    sql.NullTime   `json:"time_start"`
	TimeEnd      sql.NullTime   `json:"time_end"`
	Deadline     sql.NullTime   `json:"deadline"`
	StatusID     int64          `json:"status_id"`
}

func (q *Queries) UpdateTopic(ctx context.Context, arg UpdateTopicParams) (Topic, error) {
	row := q.db.QueryRowContext(ctx, updateTopic,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.DepartmentID,
		arg.TimeStart,
		arg.TimeEnd,
		arg.Deadline,
		arg.StatusID,
	)
	var i Topic
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.DepartmentID,
		&i.TimeStart,
		&i.TimeEnd,
		&i.Deadline,
		&i.StatusID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}