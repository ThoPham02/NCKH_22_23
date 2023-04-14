-- name: GetTopic :one
SELECT * FROM "topic"
WHERE id = $1 LIMIT 1;

-- name: ListTopics :many
SELECT * FROM "topic"
ORDER BY "name";

-- name: CreateTopic :one
INSERT INTO "topic" (
  "id", "name", "lecture_id", "faculty_id", "status", "result_url", "conference_id", "group_id", "time_start", "time_end"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: DeleteTopic :exec
DELETE FROM "topic"
WHERE id = $1;