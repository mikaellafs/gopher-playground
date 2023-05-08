package token

type TokenStore interface {
	Create(token Token) (string, error)
	Read(id string) *Token
}
