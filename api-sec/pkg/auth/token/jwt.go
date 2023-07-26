package token

import (
	"gopher-playground/api-sec/pkg/env"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewTokenClaims(username string, durationMinutes int) jwt.StandardClaims {
	now := time.Now()
	return jwt.StandardClaims{
		Subject:   username,
		Issuer:    os.Getenv(env.JWT_ISSUER),
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute * time.Duration(durationMinutes)).Unix(),
	}
}

func NewJwtToken(claims jwt.StandardClaims, signingAlgorithm, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(signingAlgorithm), claims)

	return token.SignedString([]byte(secret))
}

func ParseToken(strToken, secret string) (*jwt.StandardClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(strToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// Validate the token's claims
	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	if !ok || !parsedToken.Valid {
		log.Println("Missing claims")
		return nil, ErrInvalidToken
	}

	// Check expiration
	if claims.ExpiresAt < time.Now().Unix() {
		log.Println("Token expired")
		return nil, ErrInvalidToken
	}

	// Verify issuer
	if claims.Issuer != os.Getenv(env.JWT_ISSUER) {
		log.Println("Invalid issuer")
		return nil, ErrInvalidToken
	}

	return claims, nil
}
