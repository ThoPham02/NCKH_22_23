-- name: GetUserInfo :one
SELECT * FROM "user_info"
WHERE id = $1 LIMIT 1;

-- name: ListUserInfos :many
SELECT * FROM "user_info"
ORDER BY "name";

-- name: CreateUserInfo :one
INSERT INTO "user_info" (
  "id", "user_id", "name", "email", "phone", "faculty_id", "degree", "year_start", "avata_url", "birthday", "bank_account"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: DeleteUserInfo :exec
DELETE FROM "user_info"
WHERE id = $1;