-- name: CreateLecturer :one
INSERT INTO
    lecturers (name, user_id)
VALUES
    ($1, $2) RETURNING *;

-- name: GetLecturer :one
SELECT
    *
FROM
    lecturers
WHERE
    id = $1;

-- name: ListLecturers :many
SELECT
    *
FROM
    lecturers
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateLecturer :one
UPDATE
    lecturers
SET
    name = $2,
    updated_at = NOW()
WHERE
    id = $1 RETURNING *;

-- name: DeleteLecturer :one
UPDATE
    lecturers
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING *;