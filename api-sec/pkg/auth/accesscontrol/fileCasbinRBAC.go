package accesscontrol

import (
	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

// Implement RBAC
type FileCasbinRBAC struct {
	pathToData string
	pathToConf string
	e          *casbin.Enforcer
}

var _ AccessControl = (*FileCasbinRBAC)(nil)

func NewFileCasbinRBAC(pathToConf, pathToData string) *FileCasbinRBAC {
	return &FileCasbinRBAC{
		pathToConf: pathToConf,
		pathToData: pathToData,
	}
}

func (ac *FileCasbinRBAC) LoadPolicy() error {
	// Create an adapter to read and record policies in a file
	adapter := fileadapter.NewAdapter(ac.pathToData)

	// Create an casbin enforcer instance
	enforcer, err := casbin.NewEnforcer(ac.pathToConf, adapter)
	if err != nil {
		return err
	}

	ac.e = enforcer
	return nil
}

func (ac *FileCasbinRBAC) Enforce(values ...interface{}) (bool, error) {
	allowed, err := ac.e.Enforce(values...)
	if err != nil {
		return false, err
	}

	return allowed, nil
}

func (ac *FileCasbinRBAC) AddPolicy(p ...interface{}) error {
	_, err := ac.e.AddPolicy(p...)
	if err != nil {
		return err
	}

	return ac.e.SavePolicy()
}

func (ac *FileCasbinRBAC) AssignRoleToUser(user, role string) error {
	_, err := ac.e.AddGroupingPolicy(user, role)
	if err != nil {
		return err
	}

	return ac.e.SavePolicy()
}

func (ac *FileCasbinRBAC) RemoveUser(user string) error {
	_, err := ac.e.RemoveFilteredPolicy(0, user)
	if err != nil {
		return err
	}

	return ac.e.SavePolicy()
}
