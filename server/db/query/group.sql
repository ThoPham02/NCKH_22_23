-- name: CreateGroup :one
INSERT INTO
    groups (topic_id, lecturer_id, term_id)
VALUES
    ($1, $2, $3) RETURNING *;

-- name: GetGroup :one
SELECT
    *
FROM
    groups
WHERE
    id = $1;

-- name: ListGroups :many
SELECT
    *
FROM
    groups
ORDER BY id;

-- name: UpdateGroup :one
UPDATE
    groups
SET
    topic_id = $2,
    lecturer_id = $3,
    term_id=$4,
    updated_at = NOW()
WHERE
    id = $1 RETURNING *;

-- name: DeleteGroup :one
UPDATE
    groups
SET
    deleted_at = NOW()
WHERE
    id = $1 RETURNING *;