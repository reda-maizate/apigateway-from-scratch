-- name: CreateUser :one
INSERT INTO Users (uuid, email, password, auth_token)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM Users
WHERE email = $1;

-- name: GetUserFromAuthToken :one
SELECT *
FROM Users
WHERE auth_token = $1;