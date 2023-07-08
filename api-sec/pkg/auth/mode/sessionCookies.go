package mode

import (
	"context"
	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	sessionTokenName string = "session_token"
)

type SessionCookies struct {
	tokenStore token.TokenStore
	userRepo   user.Repository
}

var _ AuthMode = (*SessionCookies)(nil)

func NewSessionCookiesAuth(repo user.Repository, store token.TokenStore) *SessionCookies {
	return &SessionCookies{
		userRepo:   repo,
		tokenStore: store,
	}
}

func (a *SessionCookies) Authenticate(c *gin.Context) (*user.User, error) {
	// Get session token
	sTokenId, err := c.Cookie(sessionTokenName)
	if err != nil {
		return nil, err
	}

	// Read token from store
	t, err := a.tokenStore.Read(sTokenId)
	if err != nil {
		return nil, err
	}

	return a.userRepo.ReadUser(context.TODO(), t.Username)
}

func (a *SessionCookies) GenerateToken(c *gin.Context, username string, expireAt time.Time) *token.Token {
	stoken := &token.Token{
		Username:   username,
		ExpireAt:   expireAt,
		Attributes: map[string]string{},
	}

	tId, _ := a.tokenStore.Create(*stoken)

	// Add cookie
	c.SetCookie(sessionTokenName, tId, stoken.SecondsUntilExpiration(), "/", "", false, true)

	return stoken
}
