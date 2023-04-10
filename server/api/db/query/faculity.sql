-- name: GetFaculityByID :one
SELECT * FROM "faculity"
WHERE id = $1 LIMIT 1;


-- name: GetListFaculity :many
SELECT * FROM "faculity";