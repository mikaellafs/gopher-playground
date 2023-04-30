package repository

import (
	"sync"
	"time"

	"gopher-playground/api-sec/pkg/log"
)

type auditLog struct {
	method   string
	path     string
	username string
	status   int
	time     time.Time
}

type MemoryAuditLog struct {
	logs   map[int]*auditLog
	lastId int

	mutex *sync.RWMutex
}

var _ (log.Repository) = (*MemoryAuditLog)(nil)

func NewMemoryAuditLogRepository() *MemoryAuditLog {
	return &MemoryAuditLog{
		logs:  map[int]*auditLog{},
		mutex: &sync.RWMutex{},
	}
}

func (r *MemoryAuditLog) Add(method string, path string, username string, status int, auditTime time.Time) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.lastId += 1

	r.logs[r.lastId] = &auditLog{
		method:   method,
		path:     path,
		username: username,
		status:   status,
		time:     auditTime,
	}

	return r.lastId, nil
}

func (r *MemoryAuditLog) Update(id int, method string, path string, username string, status int, auditTime time.Time) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.logs[id] == nil {
		return log.ErrLogDoesNotExist
	}

	r.logs[id] = &auditLog{
		method:   method,
		path:     path,
		username: username,
		status:   status,
		time:     auditTime,
	}

	return nil
}

func (r *MemoryAuditLog) Read(id int) (*log.AuditLog, error) {
	if r.logs[id] == nil {
		return nil, log.ErrLogDoesNotExist
	}

	return &log.AuditLog{
		Id:       id,
		Method:   r.logs[id].method,
		Path:     r.logs[id].path,
		Username: r.logs[id].username,
		Status:   r.logs[id].status,
		Time:     r.logs[id].time,
	}, nil
}

func (r *MemoryAuditLog) ReadAll() []*log.AuditLog {
	var audlogs []*log.AuditLog

	for id, l := range r.logs {
		audlogs = append(audlogs, &log.AuditLog{
			Id:       id,
			Method:   l.method,
			Path:     l.path,
			Username: l.username,
			Status:   l.status,
			Time:     l.time,
		})
	}
	return audlogs
}
