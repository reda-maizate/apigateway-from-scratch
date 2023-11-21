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

-- name: GetUserByAuthToken :one
SELECT *
FROM Users
WHERE auth_token = $1;

-- name: GetUserPermissions :many
SELECT action FROM Permissions p
JOIN UserPermissions up ON p.uuid = up.permission_uuid
WHERE up.user_uuid = $1 AND p.service = $2 AND p.resource = $3;