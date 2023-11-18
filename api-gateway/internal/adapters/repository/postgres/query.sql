-- name: CreateNote :one
INSERT INTO Notes (uuid, title, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAllNotes :many
SELECT *
FROM Notes;

-- name: CreateUser :one
INSERT INTO Users (uuid, email, password, auth_token)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM Users
WHERE email = $1;