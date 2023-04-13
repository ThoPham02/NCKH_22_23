-- name: GetTopicRegistationById :one
SELECT * FROM "topic_registration"
WHERE id = $1 LIMIT 1;

-- name: GetTopicRegistationByLectureId :many
SELECT * FROM "topic_registration"
WHERE lecture_id = $1;

-- name: GetListTopicRegistation :many
SELECT * FROM "topic_registration"
WHERE name LIKE $1;

-- name: CreateTopicRegistation :one
INSERT INTO "topic_registration" (
  name, description, description_url, lecture_id, faculity_id, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateTopicRegistation :one
UPDATE "topic_registration"
  set name = $2,
  description =$3,
  description_url = $4,
  lecture_id = $5,
  faculity_id = $6
WHERE id = $1 RETURNING *;

-- name: DeleteTopicRegistation :exec
DELETE FROM "topic_registration"
WHERE id = $1;