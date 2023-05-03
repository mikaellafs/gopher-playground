package accesscontrol

import (
	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

type FileCasbin struct {
	pathToData string
	pathToConf string
	e          *casbin.Enforcer
}

var _ AccessControl = (*FileCasbin)(nil)

func NewFileCasbinAC(pathToConf, pathToData string) *FileCasbin {
	return &FileCasbin{
		pathToConf: pathToConf,
		pathToData: pathToData,
	}
}

func (ac *FileCasbin) LoadPolicy() error {
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

func (ac *FileCasbin) Enforce(values ...interface{}) (bool, error) {
	allowed, err := ac.e.Enforce(values...)
	if err != nil {
		return false, err
	}

	return allowed, nil
}
