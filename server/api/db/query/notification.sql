-- name: GetNotification :one
SELECT * FROM "notification"
WHERE id = $1 LIMIT 1;

-- name: ListNotifications :many
SELECT * FROM "notification"
ORDER BY "name";

-- name: CreateNotification :one
INSERT INTO "notification" (
  "id", "name", "url"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteNotification :exec
DELETE FROM "notification"
WHERE id = $1;

-- name: UpdateNotification :exec
UPDATE "notification"
  set name = $2,
  url = $3
WHERE "id" = $1;