-- name: GetTopicResult :one
SELECT * FROM "topic_result"
WHERE id = $1 LIMIT 1;

-- name: ListTopicResults :many
SELECT * FROM "topic_result";

-- name: CreateTopicResult :one
INSERT INTO "topic_result" (
  "id", "score", "comment", "topic_id", "supervisor_id"
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteTopicResult :exec
DELETE FROM "topic_result"
WHERE id = $1;