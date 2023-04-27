-- name: GetTopic :one
SELECT * FROM "topic"
WHERE id = $1 LIMIT 1;

-- name: ListTopics :many
SELECT * FROM "topic"
ORDER BY "name";

-- name: ListTopicsFilter :many
SELECT * FROM "topic"
WHERE name LIKE $1
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

-- name: AcceptTopic :exec
UPDATE "topic"
  set name = $2,
  lecture_id = $3,
  faculty_id = $4,
  status = $5,
  result_url = $6,
  conference_id = $7,
  group_id = $8,
  time_start = $9,
  time_end = $10
WHERE id = $1;

-- name: UpdateGroupTopic :exec
UPDATE "topic"
set  
  group_id = $2
WHERE id = $1;

-- name: UpdateStatusTopic :exec
UPDATE "topic"
set  
  status = $2
WHERE id = $1;