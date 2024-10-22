package services

import (
    "golang.org/x/crypto/bcrypt"
)

type BcryptService struct{}

func NewBcryptService() *BcryptService {
    return &BcryptService{}
}

func (b *BcryptService) HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashedPassword), err
}

func (b *BcryptService) VerifyPassword(password, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
