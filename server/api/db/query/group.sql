-- name: GetGroup :one
SELECT * FROM "group"
WHERE id = $1 LIMIT 1;

-- name: ListGroups :many
SELECT * FROM "group"
ORDER BY "type_account";

-- name: CreateGroup :one
INSERT INTO "group" (
  "id", "name", "conference_id", "faculty_id"
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM "group"
WHERE id = $1;