package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Claims struct {
	jwt.RegisteredClaims
	ID           string `json:"id"`
	Name         string `json:"name"`
	Role         []Role `json:"role"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type Jwt struct {
	logger Logger
	env    Env
}

func NewJwt(logger Logger, env Env) Jwt {
	return Jwt{
		logger: logger,
		env:    env,
	}
}

func (j Jwt) Decode(jwtString string) (*Claims, error) {

	token, err := jwt.Parse(jwtString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(j.env.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("unable get claims")
	}

	c := Claims{}
	err = mapstructure.Decode(claims, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (j Jwt) Encode(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.env.JwtSecret))
}
