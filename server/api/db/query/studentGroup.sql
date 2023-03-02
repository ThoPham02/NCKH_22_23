-- name: CreateStudentGroup :one
INSERT INTO
    student_groups (
        student_id,
        group_id
    )
VALUES
    ($1, $2 ) RETURNING *;

-- name: ListStudentGroups :many
SELECT
    *
FROM
    student_groups
WHERE
    group_id = $1 AND deleted_at IS NULL;

-- name: UpdateStudentGroup :one
UPDATE
    student_groups
SET
    student_id = $2
WHERE
    group_id = $1 RETURNING *;

-- name: DeleteStudentGroup :exec
DELETE FROM student_groups
WHERE student_id = $2 AND group_id = $1;