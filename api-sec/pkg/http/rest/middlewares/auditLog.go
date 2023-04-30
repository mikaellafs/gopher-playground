package middlewares

import (
	"log"
	"time"

	"gopher-playground/api-sec/pkg/auth/user"
	apilog "gopher-playground/api-sec/pkg/log"

	"github.com/gin-gonic/gin"
)

const (
	auditLogIdName string = "audit_log_id"
)

type AuditLog struct {
	repo apilog.Repository
}

func NewAuditLog(r apilog.Repository) *AuditLog {
	return &AuditLog{
		repo: r,
	}
}

func (m *AuditLog) StartMiddleware(c *gin.Context) {
	defer c.Next()

	// Get user from context
	username := getUsernameFromContext(c)

	// Save log with no status
	id, err := m.repo.Add(c.Request.Method, c.FullPath(), username, -1, time.Now())
	if err != nil {
		log.Println("Failed to save log to repository:", err.Error())
		return
	}

	// Add audit log id to request context
	c.Set(auditLogIdName, id)
}

func (m *AuditLog) EndMiddleware(c *gin.Context) {
	defer c.Next()

	// Get audit log id from request context
	iid, exists := c.Get(auditLogIdName)
	now := time.Now()

	// Get user from context
	username := getUsernameFromContext(c)

	var err error
	if !exists {
		_, err = m.repo.Add(c.Request.Method, c.FullPath(), username, c.Writer.Status(), now)
	} else {
		id, _ := iid.(int)
		err = m.repo.Update(id, c.Request.Method, c.FullPath(), username, c.Writer.Status(), now)
	}

	if err != nil {
		log.Println("Failed to save log to repository:", err.Error())
	}
}

func getUsernameFromContext(c *gin.Context) string {
	iu, exists := c.Get("user")

	var username string
	if exists {
		u, _ := iu.(*user.User)
		username = u.Username
	}

	return username
}
