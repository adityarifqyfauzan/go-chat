package jwt

import (
	"fmt"
	"time"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func SignToken(params *config.Params, userID int) (string, error) {
	jwtKey := []byte(params.Env.GetString("app.secret"))

	claims := CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    params.Env.GetString("app.name"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyWithClaims(params *config.Params, tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	jwtKey := []byte(params.Env.GetString("app.secret"))
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid signature")
		}
		return nil, fmt.Errorf("could not parse token: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
