package ports

type PermissionService interface {
	CheckPermission(UserUuid, Service, Resource, Action string) (bool, error)
}

type PermissionRepository interface {
	CheckPermission(UserUuid, Service, Resource, Action string) (bool, error)
}
