package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"gopher-playground/api-sec/pkg/env"
	"os"

	"golang.org/x/crypto/scrypt"
)

func Encrypt(password string) (string, error) {
	// Get salt
	salt := os.Getenv(env.SALT_PASSWORD_ENCRYPT)

	dk, err := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	// Encode password to base64
	h := sha1.New()
	h.Write(dk)

	return hex.EncodeToString(h.Sum(nil)), nil
}

func CompareEqual(hash, password string) bool {
	hashedPwd, err := Encrypt(password)
	if err != nil {
		return false
	}

	return hash == hashedPwd
}
