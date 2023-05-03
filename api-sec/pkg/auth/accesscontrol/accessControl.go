package accesscontrol

type AccessControl interface {
	LoadPolicy() error

	Enforce(values ...interface{}) (bool, error)

	// AddPolicy(p string) error
	// RemovePolicy(p string) error
}
