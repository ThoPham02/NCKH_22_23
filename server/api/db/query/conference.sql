-- name: GetConference :one
SELECT * FROM "conference"
WHERE id = $1 LIMIT 1;

-- name: ListConferences :many
SELECT * FROM "conference"
ORDER BY "name";

-- name: CreateConference :one
INSERT INTO "conference" (
  "id", "name", "cash_support", "school_year"
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteConference :exec
DELETE FROM "conference"
WHERE id = $1;