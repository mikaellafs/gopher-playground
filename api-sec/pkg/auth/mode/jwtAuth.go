package mode

import (
	"context"
	"os"
	"strings"
	"time"

	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"
	"gopher-playground/api-sec/pkg/env"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type JwtAuth struct {
	userRepo                                                   user.Repository
	signingAlgorithm                                           string
	durationMinutes, refreshDurationMinutes, maxRefreshMinutes int
}

var _ AuthMode = (*JwtAuth)(nil)

func NewJwtAuth(repo user.Repository, signingAlgorithm string, durationMinutes, refreshDurationMinutes, maxRefreshMinutes int) *JwtAuth {
	return &JwtAuth{
		userRepo:               repo,
		signingAlgorithm:       signingAlgorithm,
		durationMinutes:        durationMinutes,
		refreshDurationMinutes: refreshDurationMinutes,
		maxRefreshMinutes:      maxRefreshMinutes,
	}
}

func (m *JwtAuth) Authenticate(c *gin.Context) (*user.User, error) {
	// Get token from header
	authorizationHeader := c.GetHeader("Authorization")

	// Check if the header is not empty and starts with "Bearer "
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return nil, ErrInvalidAuthHeader
	}

	// Extract the access token from the "Authorization" header
	accessToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// Parse token
	claims, err := token.ParseToken(accessToken, os.Getenv(env.JWT_SECRET_ACC))
	if err != nil {
		return nil, err
	}

	// Get user
	return m.userRepo.ReadUser(context.TODO(), claims.Subject)
}

func (m *JwtAuth) GenerateToken(c *gin.Context, username string) (*token.Token, error) {
	// Generate access token
	accessClaims := token.NewTokenClaims(username, m.durationMinutes)
	accessToken, err := token.NewJwtToken(accessClaims, m.signingAlgorithm, os.Getenv(env.JWT_SECRET_ACC))
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate auth token")
	}
	// Generate refresh token
	refreshClaims := token.NewTokenClaims(username, m.refreshDurationMinutes)
	refreshToken, err := token.NewJwtToken(accessClaims, m.signingAlgorithm, os.Getenv(env.JWT_SECRET_REFR))
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate refresh token")
	}

	// Add set cookie with tokens
	c.SetCookie("access_token", accessToken, int(accessClaims.ExpiresAt), "/", "", true, true)
	c.SetCookie("refresh_token", refreshToken, int(refreshClaims.ExpiresAt), "/", "", true, true)

	// Create token to return
	return &token.Token{
		Username:   username,
		ExpireAt:   time.Unix(accessClaims.ExpiresAt, 0),
		Attributes: map[string]string{},
	}, nil
}

func (m *JwtAuth) Logout(c *gin.Context) {

}
