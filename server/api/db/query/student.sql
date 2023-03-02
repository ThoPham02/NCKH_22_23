-- name: CreateStudent :one
INSERT INTO
    students (name, user_id)
VALUES
    ($1, $2) RETURNING *;

-- name: GetStudent :one
SELECT
    *
FROM
    students
WHERE
    id = $1;

-- name: ListStudents :many
SELECT
    *
FROM
    students
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateStudent :one
UPDATE
    students
SET
    name = $2,
    updated_at = NOW()
WHERE
    id = $1 RETURNING *;

-- name: DeleteStudent :one
UPDATE
    students
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING *;