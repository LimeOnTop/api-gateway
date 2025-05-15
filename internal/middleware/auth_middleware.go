package middleware

import (
	"net/http"
	"strings"

	"api-gateway/gen/auth"
	"api-gateway/internal/client"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authClient *client.AuthClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		refreshToken := ctx.GetHeader("Refresh-Token")
		if accessToken == authHeader {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not find bearer token in authorization header"})
			return
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		// Проверяем токен через auth service
		res, err := authClient.ValidateToken(ctx.Request.Context(), &auth.ValidateTokenRequest{AccessToken: accessToken})
		if err != nil || !res.Valid {
			// Пробуем обновить токен
			newRefreshToken, refreshErr := authClient.RefreshToken(ctx.Request.Context(), &auth.RefreshTokenRequest{RefreshToken: refreshToken})
			if refreshErr != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				return
			}
			// Устанавливаем новые токены в заголовки ответа
			ctx.Header("New-Access-Token", newRefreshToken.AccessToken)
		}

	}
}
