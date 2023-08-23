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

const (
	refreshTokenName string = "refresh_token"
)

type JwtAuth struct {
	userRepo                                user.Repository
	signingAlgorithm                        string
	durationMinutes, refreshDurationMinutes int
}

var _ AuthMode = (*JwtAuth)(nil)

func NewJwtAuth(repo user.Repository, signingAlgorithm string, durationMinutes, refreshDurationMinutes int) *JwtAuth {
	return &JwtAuth{
		userRepo:               repo,
		signingAlgorithm:       signingAlgorithm,
		durationMinutes:        durationMinutes,
		refreshDurationMinutes: refreshDurationMinutes,
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
	refreshToken, err := token.NewJwtToken(refreshClaims, m.signingAlgorithm, os.Getenv(env.JWT_SECRET_REFR))
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate refresh token")
	}

	// Add auth header
	c.Header("Authorization", "Bearer "+accessToken)

	// Add set cookie header for refresh token
	c.SetCookie(refreshTokenName, refreshToken, int(refreshClaims.ExpiresAt), "/", "", true, true)

	// Create token to return
	return &token.Token{
		Username:   username,
		ExpireAt:   time.Unix(accessClaims.ExpiresAt, 0),
		Attributes: map[string]string{},
	}, nil
}

func (m *JwtAuth) Logout(c *gin.Context) {
	// Could implement a black list and add token to it, however, lets keep it simple
	// Remove token from cookies and header
	c.Header("Authorization", "")
	c.SetCookie(refreshTokenName, "", -1, "/", "", true, true)
}

func (m *JwtAuth) Refresh(c *gin.Context) (*token.Token, error) {
	// Get refresh token from cookies
	strToken, err := c.Cookie(refreshTokenName)
	if err != nil {
		return nil, err
	}

	// Parse token
	parsedToken, err := token.ParseToken(strToken, os.Getenv(env.JWT_SECRET_REFR))
	if err != nil {
		return nil, err
	}

	// Generate access token
	accessClaims := token.NewTokenClaims(parsedToken.Subject, m.durationMinutes)
	accessToken, err := token.NewJwtToken(accessClaims, m.signingAlgorithm, os.Getenv(env.JWT_SECRET_ACC))
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate auth token")
	}

	// Renew auth header
	c.Header("Authorization", "Bearer "+accessToken)

	// Generate new token
	return &token.Token{
		Username:   parsedToken.Subject,
		ExpireAt:   time.Unix(accessClaims.ExpiresAt, 0),
		Attributes: map[string]string{},
	}, nil
}
