package controller

import (
	"net/http"

	"api-gateway/gen/gpt"
	"api-gateway/gen/user"
	"api-gateway/internal/client"

	"github.com/gin-gonic/gin"
)

type GptController struct {
	gptClient  *client.GptClient
	userClient *client.UserClient
}

func NewGptController(gptClient *client.GptClient, userClient *client.UserClient) *GptController {
	return &GptController{
		gptClient:  gptClient,
		userClient: userClient,
	}
}

func (c *GptController) GetGPTRecommendation(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")
	if accessToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}
	userProducts, err := c.userClient.GetUserProducts(ctx.Request.Context(), &user.UserRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user products"})
		return
	}
	userPreference, err := c.userClient.GetUserPreference(ctx.Request.Context(), &user.UserRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user preference"})
		return
	}
	// Вызываем метод GetGPTRecommendation у gptClient
	resp, err := c.gptClient.GetGPTRecommendation(ctx, &gpt.UserRequest{
		Products:   userProducts.ProductNames,
		Preference: userPreference.PreferenceName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommendation"})
		return
	}

	// Возвращаем ответ клиенту
	ctx.JSON(http.StatusOK, gin.H{
		"Message": resp.Message,
		"Image":   resp.ImageData,
		"Format":  resp.ImageFormat,
	})
}
