package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Отсутсвует заголовок авторизации"})
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(authHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("IimULHg9FRS0XleGnPZo"), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": fmt.Sprintf("Ошибка верификации токена: %s", err.Error())})
			return
		}

		ctx.Set("tgId", claims["tgId"].(string))
		ctx.Next()
	}
}
