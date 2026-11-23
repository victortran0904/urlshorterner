package service

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"

type GenPass interface {
	GeneratePassword(length int) (string, error)
}

type genPass struct {
}

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
