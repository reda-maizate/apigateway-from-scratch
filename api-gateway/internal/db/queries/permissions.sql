-- name: GetUserPermissions :many
SELECT action FROM Permissions p
JOIN UserPermissions up ON p.uuid = up.permission_uuid
WHERE up.user_uuid = $1 AND p.service = $2 AND p.resource = $3;