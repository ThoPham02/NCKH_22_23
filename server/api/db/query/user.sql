-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM "user"
WHERE name = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY "type_account";

-- name: CreateUser :one
INSERT INTO "user" (
  "id", "name", "email", "hash_password", "type_account"
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;