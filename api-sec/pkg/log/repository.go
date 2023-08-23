package log

import (
	"errors"
	"time"
)

var (
	ErrLogDoesNotExist = errors.New("Log does not exist")
)

type Repository interface {
	Add(method string, path string, username string, status int, auditTime time.Time) (int, error)
	Update(id int, method string, path string, username string, status int, auditTime time.Time) error

	Read(id int) (*AuditLog, error)
	ReadAll() []*AuditLog
}
