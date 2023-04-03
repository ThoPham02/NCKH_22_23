-- name: GetUserByName :one
SELECT * FROM "user"
WHERE name = $1 LIMIT 1;

-- name: CreateUser :exec
INSERT INTO "user" (
  name, hash_password, type_account_id, email
) VALUES (
  $1, $2, $3, $4
);