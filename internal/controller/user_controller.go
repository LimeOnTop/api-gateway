package controller

import (
	"net/http"
	"strings"

	"api-gateway/gen/user"
	"api-gateway/internal/client"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userClient *client.UserClient
}

func NewUserController(userClient *client.UserClient) *UserController {
	return &UserController{
		userClient: userClient,
	}
}

func (c *UserController) GetUserProducts(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	// Вызываем метод GetUserProducts у userClient
	resp, err := c.userClient.GetUserProducts(ctx.Request.Context(), &user.UserRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user products"})
		return
	}

	// Возвращаем ответ клиенту
	ctx.JSON(http.StatusOK, gin.H{
		"Products": resp.ProductNames,
	})
}

func (c *UserController) GetUserPreference(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	// Вызываем метод GetUserPreference у userClient
	resp, err := c.userClient.GetUserPreference(ctx.Request.Context(), &user.UserRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user preference"})
		return
	}

	// Возвращаем ответ клиенту
	ctx.JSON(http.StatusOK, gin.H{
		"Preference": resp.PreferenceName,
	})
}

func (c *UserController) AddProduct(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	var req user.AddProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	req.AccessToken = accessToken

	// Вызываем метод AddUserProduct у userClient
	resp, err := c.userClient.AddUserProduct(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) RemoveProduct(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	var req user.RemoveProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	req.AccessToken = accessToken

	// Вызываем метод RemoveUserProduct у userClient
	resp, err := c.userClient.RemoveUserProduct(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove product"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) UpdatePreference(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	var req user.UpdatePreferenceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	req.AccessToken = accessToken

	resp, err := c.userClient.UpdateUserPreference(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update preference"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) RemovePreference(ctx *gin.Context) {
	accessToken := c.getAuthHeader(ctx)

	var req user.RemovePreferenceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	req.AccessToken = accessToken

	resp, err := c.userClient.RemoveUserPreference(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove preference"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) getAuthHeader(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return ""
	}
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	return accessToken
}
