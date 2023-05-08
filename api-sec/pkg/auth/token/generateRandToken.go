package token

import (
	"crypto/rand"
	"encoding/base64"
)

func generateRandToken(n int) ([]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)

	return token, err
}

func GenerateRandTokenString() (string, error) {
	token, err := generateRandToken(32)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(token), nil
}
