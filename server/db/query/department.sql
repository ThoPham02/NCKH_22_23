-- name: CreateDepartment :one
INSERT INTO
    departments (name, faculty_id, user_id)
VALUES
    ($1, $2, $3) RETURNING *;

-- name: GetDepartment :one
SELECT
    *
FROM
    departments
WHERE
    id = $1;

-- name: ListDepartments :many
SELECT
    *
FROM
    departments
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateDepartment :one
UPDATE
    departments
SET
    name = $2,
    faculty_id = $3,
    updated_at = NOW()
WHERE
    id = $1 RETURNING *;

-- name: DeleteDepartment :one
UPDATE
    departments
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING *;