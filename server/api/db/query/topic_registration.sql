-- name: GetTopicRegistration :one
SELECT * FROM "topic_registration"
WHERE id = $1 LIMIT 1;

-- name: ListTopicRegistrations :many
SELECT * FROM "topic_registration"
WHERE name like $1 
ORDER BY "name";

-- name: CreateTopicRegistration :one
INSERT INTO "topic_registration" (
  "id", "name", "lecture_id", "faculty_id", "status", "created_at"
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteTopicRegistration :exec
DELETE FROM "topic_registration"
WHERE id = $1;

-- name: UpdateTopicRegistration :exec
UPDATE "topic_registration"
  set status = $2
WHERE "id" in ($1);