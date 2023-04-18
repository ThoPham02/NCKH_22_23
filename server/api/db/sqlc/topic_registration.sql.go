// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: topic_registration.sql

package db

import (
	"context"
	"time"
)

const createTopicRegistration = `-- name: CreateTopicRegistration :one
INSERT INTO "topic_registration" (
  "id", "name", "lecture_id", "faculty_id", "status", "created_at"
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, name, lecture_id, faculty_id, status, created_at
`

type CreateTopicRegistrationParams struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	LectureID int32     `json:"lecture_id"`
	FacultyID int32     `json:"faculty_id"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateTopicRegistration(ctx context.Context, arg CreateTopicRegistrationParams) (TopicRegistration, error) {
	row := q.db.QueryRowContext(ctx, createTopicRegistration,
		arg.ID,
		arg.Name,
		arg.LectureID,
		arg.FacultyID,
		arg.Status,
		arg.CreatedAt,
	)
	var i TopicRegistration
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LectureID,
		&i.FacultyID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTopicRegistration = `-- name: DeleteTopicRegistration :exec
DELETE FROM "topic_registration"
WHERE id = $1
`

func (q *Queries) DeleteTopicRegistration(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTopicRegistration, id)
	return err
}

const getTopicRegistration = `-- name: GetTopicRegistration :one
SELECT id, name, lecture_id, faculty_id, status, created_at FROM "topic_registration"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTopicRegistration(ctx context.Context, id int32) (TopicRegistration, error) {
	row := q.db.QueryRowContext(ctx, getTopicRegistration, id)
	var i TopicRegistration
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LectureID,
		&i.FacultyID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listTopicRegistrations = `-- name: ListTopicRegistrations :many
SELECT id, name, lecture_id, faculty_id, status, created_at FROM "topic_registration"
WHERE name like $1 
ORDER BY "name"
`

func (q *Queries) ListTopicRegistrations(ctx context.Context, name string) ([]TopicRegistration, error) {
	rows, err := q.db.QueryContext(ctx, listTopicRegistrations, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TopicRegistration{}
	for rows.Next() {
		var i TopicRegistration
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.LectureID,
			&i.FacultyID,
			&i.Status,
			&i.CreatedAt,
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

const updateTopicRegistration = `-- name: UpdateTopicRegistration :exec
UPDATE "topic_registration"
  set status = $2
WHERE "id" in ($1)
`

type UpdateTopicRegistrationParams struct {
	ID     int32 `json:"id"`
	Status int32 `json:"status"`
}

func (q *Queries) UpdateTopicRegistration(ctx context.Context, arg UpdateTopicRegistrationParams) error {
	_, err := q.db.ExecContext(ctx, updateTopicRegistration, arg.ID, arg.Status)
	return err
}
