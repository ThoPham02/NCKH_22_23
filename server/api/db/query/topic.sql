-- name: CreateTopic :one
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
    ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetTopic :one
SELECT
    *
FROM
    topics
WHERE
    deleted_at IS NULL
    AND id = $1;

-- name: ListTopics :many
SELECT
    *
FROM
    topics
WHERE
    deleted_at IS NULL
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateTopic :one
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
    id = $1 RETURNING *;

-- name: DeleteTopic :exec
Update
    topics
SET
    deleted_at = NOW()
WHERE
    id = $1;