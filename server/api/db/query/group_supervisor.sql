-- name: GetGroupSupervisor :one
SELECT * FROM "group_supervisor"
WHERE id = $1 LIMIT 1;

-- name: ListGroupSupervisors :many
SELECT * FROM "group_supervisor";

-- name: CreateGroupSupervisor :one
INSERT INTO "group_supervisor" (
  "id", "supervisor_id", "role", "group_report_id"
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteGroupSupervisor :exec
DELETE FROM "group_supervisor"
WHERE id = $1;