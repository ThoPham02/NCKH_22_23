-- name: CreateUser :one
INSERT INTO
    users (username, password, email, permission)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT
    *
FROM
    users
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE
    users
SET
    username = $2,
    password = $3,
    email=$4,
    permission = $5,
    updated_at = NOW()
WHERE
    id = $1 RETURNING *;

-- name: DeleteUser :one
UPDATE
    users
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING *;