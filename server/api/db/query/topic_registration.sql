-- name: GetTopicRegistration :one
SELECT * FROM "topic_registration"
WHERE id = $1 LIMIT 1;

-- name: ListGetTopicRegistrations :many
SELECT * FROM "topic_registration"
ORDER BY "name";

-- name: CreateGetTopicRegistration :one
INSERT INTO "topic_registration" (
  "id", "name", "lecture_id", "faculty_id", "status", "created_at"
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteGetTopicRegistration :exec
DELETE FROM "topic_registration"
WHERE id = $1;