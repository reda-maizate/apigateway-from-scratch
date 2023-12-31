// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: permissions.sql

package db

import (
	"context"
)

const getUserPermissions = `-- name: GetUserPermissions :many
SELECT action FROM Permissions p
JOIN UserPermissions up ON p.uuid = up.permission_uuid
WHERE up.user_uuid = $1 AND p.service = $2 AND p.resource = $3
`

type GetUserPermissionsParams struct {
	UserUuid string
	Service  string
	Resource string
}

func (q *Queries) GetUserPermissions(ctx context.Context, arg GetUserPermissionsParams) ([]string, error) {
	rows, err := q.db.Query(ctx, getUserPermissions, arg.UserUuid, arg.Service, arg.Resource)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var action string
		if err := rows.Scan(&action); err != nil {
			return nil, err
		}
		items = append(items, action)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
