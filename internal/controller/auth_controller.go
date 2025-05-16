package controller

import (
	"net/http"
	"strings"

	"api-gateway/gen/auth"
	"api-gateway/internal/client"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authClient *client.AuthClient
}

func NewAuthController(authClient *client.AuthClient) *AuthController {
	return &AuthController{
		authClient: authClient,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	resp, err := c.authClient.Register(ctx.Request.Context(), &auth.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": resp.UserId,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	resp, err := c.authClient.Login(ctx.Request.Context(), &auth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
	})
	ctx.Header("Authorization", resp.AccessToken)
	ctx.Header("Refresh-Token", resp.RefreshToken)
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("Refresh-Token")
	if refreshToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is required"})
		return
	}

	resp, err := c.authClient.RefreshToken(ctx.Request.Context(), &auth.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": resp.AccessToken,
	})
	ctx.Header("Authorization", resp.AccessToken)
}

func (c *AuthController) ValidateToken(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	resp, err := c.authClient.ValidateToken(ctx.Request.Context(), &auth.ValidateTokenRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"valid":   resp.Valid,
		"user_id": resp.UserId,
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	resp, err := c.authClient.Logout(ctx.Request.Context(), &auth.LogoutRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
	ctx.Header("Authorization", "")
	ctx.Header("Refresh-Token", "")
}

func (c *AuthController) getAuthHeader(ctx *gin.Context) string {
	authHeader:= ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return ""
	}
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	return accessToken
}
