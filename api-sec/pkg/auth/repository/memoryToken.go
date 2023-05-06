package repository

import (
	"sync"
	"time"

	"gopher-playground/api-sec/pkg/auth/token"

	"github.com/google/uuid"
)

type mtoken struct {
	id       string
	username string
	expireAt time.Time
	att      map[string]string
}

type MemoryToken struct {
	tokens map[string]*mtoken
	mutex  *sync.Mutex
}

var _ token.TokenStore = (*MemoryToken)(nil)

func NewInMemoryTokenStore() *MemoryToken {
	return &MemoryToken{
		tokens: map[string]*mtoken{},
		mutex:  &sync.Mutex{},
	}
}

func (r *MemoryToken) Create(t token.Token) string {
	newId := uuid.New().String()

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.tokens[newId] = &mtoken{
		id:       newId,
		username: t.Username,
		expireAt: t.ExpireAt,
		att:      t.Attributes,
	}

	return newId
}

func (r *MemoryToken) Read(id string) *token.Token {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.tokens[id] == nil {
		return nil
	}

	return &token.Token{
		Username:   r.tokens[id].username,
		ExpireAt:   r.tokens[id].expireAt,
		Attributes: r.tokens[id].att,
	}
}

func (r *MemoryToken) Delete(id string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.tokens[id] = nil
}
