-- name: CreateNote :one
INSERT INTO Notes (uuid, title, content, created_by)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllNotes :many
SELECT *
FROM Notes;