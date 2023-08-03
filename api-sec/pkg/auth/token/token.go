package token

import (
	"time"
)

type Token struct {
	Username   string
	ExpireAt   time.Time
	Attributes map[string]string
}

func (t *Token) TimeUntilExpiration() time.Duration {
	return t.ExpireAt.Sub(time.Now())
}

func (t *Token) SecondsUntilExpiration() int {
	timeLeft := t.TimeUntilExpiration()

	return int(timeLeft.Seconds())
}
