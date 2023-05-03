package accesscontrol

type AccessControl interface {
	LoadPolicy() error

	Enforce(values ...interface{}) (bool, error)

	AddPolicy(p ...interface{}) error
	// RemovePolicy(p string) error

	AssignRoleToUser(user, role string) error
	RemoveUser(user string) error
}
