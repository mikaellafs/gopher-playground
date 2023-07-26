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

	// Check token is expired
	if t.SecondsUntilExpiration() <= 0 {
		a.tokenStore.Delete(sTokenId)
		return nil, token.ErrInvalidToken
	}

	return a.userRepo.ReadUser(context.TODO(), t.Username)
}

func (a *SessionCookies) GenerateToken(c *gin.Context, username string) (*token.Token, error) {
	stoken := &token.Token{
		Username:   username,
		ExpireAt:   time.Now().Add(30 * time.Minute),
		Attributes: map[string]string{},
	}

	tId, err := a.tokenStore.Create(*stoken)
	if err != nil {
		return nil, err
	}

	// Add cookie
	c.SetCookie(sessionTokenName, tId, stoken.SecondsUntilExpiration(), "/", "", false, true)

	return stoken, nil
}

func (a *SessionCookies) Logout(c *gin.Context) {
	sTokenId, err := c.Cookie(sessionTokenName)
	if err != nil {
		return
	}

	// Delete from store
	err = a.tokenStore.Delete(sTokenId)
	if err != nil {
		return
	}

	// Clean cookie
	c.SetCookie(sessionTokenName, "", -1, "", "", false, true)
}
