package jwt

import (
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
