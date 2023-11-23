package ports

type CheckPermissionParams struct {
	UserUuid string
	Service  string
	Resource string
	Action   string
}

type CheckPermissionResponse struct {
	Authorized bool
}

type PermissionService interface {
	CheckPermission(CheckPermissionParams) (CheckPermissionResponse, error)
}

type PermissionRepository interface {
	CheckPermission(CheckPermissionParams) (CheckPermissionResponse, error)
}
