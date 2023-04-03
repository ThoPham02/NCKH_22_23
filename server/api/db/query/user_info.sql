-- name: GetUserInfo :one
SELECT * FROM "user_info"
WHERE user_id = $1 LIMIT 1;