-- name: GetUserInfo :one
SELECT * FROM "user_info"
WHERE user_id = $1 LIMIT 1;

-- name: CreateUserInfo :exec
INSERT INTO "user_info" (
  user_id, name, description, avata_url, birthday, faculity_id, year_start, bank_account, phone, sex
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
);

-- name: UpdateUserInfo :exec
UPDATE "user_info"
  set name = $2,
  description = $3,
  avata_url = $4,
  birthday = $5,
  faculity_id = $6,
  year_start = $7,
  bank_account = $8,
  phone = $9,
  sex = $10
WHERE user_id = $1;