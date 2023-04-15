-- name: GetUserInfo :one
SELECT * FROM "user_info"
WHERE "user_id" = $1 LIMIT 1;

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

-- name: UpdateUserInfo :exec
UPDATE "user_info"
  set name = $2,
  email = $3,
  phone = $4,
  faculty_id = $5,
  degree = $6,
  year_start = $7,
  avata_url = $8,
  birthday = $9,
  bank_account = $10
WHERE "user_id" = $1;

-- name: ListUserInfosByType :many
SELECT * FROM "user_info"
WHERE user_id in (SELECT id FROM "user" WHERE "type_account" = $1)
ORDER BY "name";