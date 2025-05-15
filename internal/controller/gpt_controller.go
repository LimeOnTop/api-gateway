package controller

import (
	// "net/http"

	"api-gateway/internal/client"
	// "github.com/gin-gonic/gin"
)

type GptController struct {
	gptClient *client.GptClient
}

func NewGptClient(gptClient *client.GptClient) *GptController {
	return &GptController{
		gptClient: gptClient,
	}
}

