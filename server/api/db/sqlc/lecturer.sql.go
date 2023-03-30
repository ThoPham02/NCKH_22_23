// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: lecturer.sql

package db

import (
	"context"
)

const createLecturer = `-- name: CreateLecturer :one
INSERT INTO
    lecturers (name, user_id)
VALUES
    ($1, $2) RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

type CreateLecturerParams struct {
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
}

func (q *Queries) CreateLecturer(ctx context.Context, arg CreateLecturerParams) (Lecturer, error) {
	row := q.db.QueryRowContext(ctx, createLecturer, arg.Name, arg.UserID)
	var i Lecturer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteLecturer = `-- name: DeleteLecturer :one
UPDATE
    lecturers
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteLecturer(ctx context.Context, id int64) (Lecturer, error) {
	row := q.db.QueryRowContext(ctx, deleteLecturer, id)
	var i Lecturer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getLecturer = `-- name: GetLecturer :one
SELECT
    id, name, user_id, created_at, updated_at, deleted_at
FROM
    lecturers
WHERE
    user_id = $1
`

func (q *Queries) GetLecturer(ctx context.Context, userID int64) (Lecturer, error) {
	row := q.db.QueryRowContext(ctx, getLecturer, userID)
	var i Lecturer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listLecturers = `-- name: ListLecturers :many
SELECT
    id, name, user_id, created_at, updated_at, deleted_at
FROM
    lecturers
ORDER BY
    id
LIMIT $1 OFFSET $2
`

type ListLecturersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListLecturers(ctx context.Context, arg ListLecturersParams) ([]Lecturer, error) {
	rows, err := q.db.QueryContext(ctx, listLecturers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Lecturer{}
	for rows.Next() {
		var i Lecturer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
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

const updateLecturer = `-- name: UpdateLecturer :one
UPDATE
    lecturers
SET
    name = $2,
    updated_at = NOW()
WHERE
    id = $1 RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

type UpdateLecturerParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateLecturer(ctx context.Context, arg UpdateLecturerParams) (Lecturer, error) {
	row := q.db.QueryRowContext(ctx, updateLecturer, arg.ID, arg.Name)
	var i Lecturer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}