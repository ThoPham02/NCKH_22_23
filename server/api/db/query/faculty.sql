-- name: GetFaculty :one
SELECT * FROM "faculty"
WHERE id = $1 LIMIT 1;

-- name: ListFaculties :many
SELECT * FROM "faculty"
ORDER BY "name";

-- name: CreateFaculty :one
INSERT INTO "faculty" (
  "id", "name"
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteFaculty :exec
DELETE FROM "faculty"
WHERE id = $1;