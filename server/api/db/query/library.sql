-- name: GetLibrary :one
SELECT * FROM "library"
WHERE id = $1 LIMIT 1;

-- name: ListLibrarys :many
SELECT * FROM "library"
ORDER BY "name";

-- name: CreateLibrary :one
INSERT INTO "library" (
  "id", "name", "url", "owner_id"
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteLibrary :exec
DELETE FROM "library"
WHERE id = $1;