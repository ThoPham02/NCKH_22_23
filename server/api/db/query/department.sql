-- name: GetDepartment :one
SELECT * FROM "department"
WHERE id = $1 LIMIT 1;

-- name: ListDepartments :many
SELECT * FROM "department"
ORDER BY "type_account";

-- name: CreateDepartment :one
INSERT INTO "department" (
  "id", "name", "faculty_id"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteDepartment :exec
DELETE FROM "department"
WHERE id = $1;

-- name: UpdateDepartment :exec
UPDATE "department"
  set name = $2,
  faculty_id = $3
WHERE "id" = $1;