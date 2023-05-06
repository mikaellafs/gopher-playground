package token

import (
	"time"
)

type Token struct {
	Username   string
	ExpireAt   time.Time
	Attributes map[string]string
}
