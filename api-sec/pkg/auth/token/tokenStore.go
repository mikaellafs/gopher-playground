package token

type TokenStore interface {
	Create(token Token) string
	Read(id string) *Token
}
