package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type (
	Service struct {
		conf Config
	}

	Claims struct {
		jwt.StandardClaims
		UserID  int    `json:"userId"`
		Address string `json:"address"`
		Email   string `json:"email"`
	}
)

func (s *Service) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.conf.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, errors.New("Unable to parse token")
}

func NewService(conf Config) *Service {
	return &Service{conf: conf}
}
