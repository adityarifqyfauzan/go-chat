package middleware

import (
	"strings"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/pkg/exception"
	"github.com/adityarifqyfauzan/go-chat/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(params *config.Params) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			panic(exception.NewCredentialException("Authorization header is missing"))
		}

		// Split the header to get the Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			panic(exception.NewCredentialException("Invalid Authorization header format"))
		}

		token := parts[1]
		claims, err := jwt.VerifyWithClaims(params, token)
		if err != nil {
			panic(exception.NewCredentialException(err.Error()))
		}

		// set authenticated user id
		ctx.Set("user_id", claims.UserID)

		ctx.Next()
	}
}
