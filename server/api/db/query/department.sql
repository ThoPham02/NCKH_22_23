-- name: GetDepartmentByID :one
SELECT * FROM "department"
WHERE id = $1 LIMIT 1;


-- name: GetListDepartment :many
SELECT * FROM "department";

-- name: GetListDepartmentByFaculity :many
SELECT * FROM "department"
WHERE faculity_id = $1;