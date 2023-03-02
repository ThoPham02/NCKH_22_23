-- name: CreateReference :one
INSERT INTO
    "references" (reference_url, name)
VALUES
    ($1, $2) RETURNING *;

-- name: GetReference :one
SELECT
    *
FROM
    "references"
WHERE
    deleted_at IS NULL
    AND id = $1;

-- name: ListReferences :many
SELECT
    *
FROM
    "references"
WHERE
    deleted_at IS NULL
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateReference :one
UPDATE
    "references"
SET
    reference_url = $2,
    name = $3
WHERE
    id = $1 RETURNING *;

-- name: DeleteReference :exec
Update
    "references"
SET
    deleted_at = NOW()
WHERE
    id = $1;