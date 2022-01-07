package service

import (
	"time"

	"github.com/emregocer/golang_project/config"
	model "github.com/emregocer/golang_project/internal/model/jwt"
	"github.com/golang-jwt/jwt"
)

type JwtService struct {
	config config.Config
}

func NewJwtService(config config.Config) *JwtService {
	return &JwtService{config: config}
}

func (j *JwtService) GenerateToken(userId int) (*model.Claims, string, error) {
	t := time.Now()
	claim := &model.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "golang_api",
			ExpiresAt: t.Add(60 * time.Minute).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
		UserID: userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(j.config.JWTKey))

	return claim, tokenString, err
}
