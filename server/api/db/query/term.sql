-- name: CreateTerm :one
INSERT INTO
    terms (
        name
    )
VALUES
    ($1 ) RETURNING *;

-- name: GetTerm :one
SELECT
    *
FROM
    terms
WHERE
    deleted_at IS NULL
    AND id = $1;

-- name: ListTerms :many
SELECT
    *
FROM
    terms
WHERE
    deleted_at IS NULL;

-- name: UpdateTerm :one
UPDATE
    terms
SET
    name = $2
WHERE
    id = $1 RETURNING *;

-- name: DeleteTerm :exec
Update
    terms
SET
    deleted_at = NOW()
WHERE
    id = $1;