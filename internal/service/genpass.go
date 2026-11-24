package service

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"

//go:generate mockery --name=GenPass --filename genpass.go --outpkg mocks

// GenPass generates passwords with cryptographically secure randomness.
type GenPass interface {
	// GeneratePassword returns a random password with the requested number of ASCII characters.
	// The length must be nonnegative. Errors from the random source are returned.
	GeneratePassword(length int) (string, error)
}

type genPass struct {
}

// NewGenPass creates a generator for passwords drawn from ASCII letters, digits, and symbols.
func NewGenPass() GenPass {
	return &genPass{}
}

func (g *genPass) GeneratePassword(length int) (string, error) {
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}

	return string(password), nil
}
