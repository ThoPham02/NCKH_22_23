-- name: GetStudentTopic :one
SELECT * FROM "student_topic"
WHERE id = $1 LIMIT 1;

-- name: ListStudentTopics :many
SELECT * FROM "student_topic"
ORDER BY "type_account";

-- name: CreateStudentTopic :one
INSERT INTO "student_topic" (
  "id", "student_id", "topic_id"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteStudentTopic :exec
DELETE FROM "student_topic"
WHERE id = $1;

-- name: UpdateStudentTopic :exec
UPDATE "student_topic"
  set student_id = $2,
  topic_id = $3
WHERE "id" = $1;