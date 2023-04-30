package log

import "time"

type AuditLog struct {
	Id       int       `json:"id"`
	Method   string    `json:"method"`
	Path     string    `json:"path"`
	Username string    `json:"username"`
	Status   int       `json:"status"`
	Time     time.Time `json:"audit_time"`
}
