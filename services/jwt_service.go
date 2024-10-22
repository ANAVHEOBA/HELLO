package services

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

type JwtService struct {
    Secret string
}

type Claims struct {
    UserID string `json:"sub"`
    Exp    int64  `json:"exp"`
}

func NewJwtService(secret string) *JwtService {
    return &JwtService{Secret: secret}
}

func (j *JwtService) CreateToken(userID string) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour).Unix()

    claims := Claims{
        UserID: userID,
        Exp:    expirationTime,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.Secret))
}

func (j *JwtService) ValidateToken(tokenString string) (string, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(j.Secret), nil
    })

    if err != nil || !token.Valid {
        return "", err
    }

    return claims.UserID, nil
}
