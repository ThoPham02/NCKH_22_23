-- name: CreateFaculty :one
INSERT INTO
    faculties (name, user_id)
VALUES
    ($1, $2) RETURNING *;

-- name: GetFaculty :one
SELECT
    *
FROM
    faculties
WHERE
    id = $1;

-- name: ListFaculties :many
SELECT
    *
FROM
    faculties
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateFaculty :one
UPDATE
    faculties
SET
    name = $2,
    updated_at = NOW()
WHERE
    id = $1 RETURNING *;

-- name: DeleteFaculty :one
UPDATE
    faculties
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING *;