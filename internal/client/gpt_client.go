package client

import (
	"context"
	"log"

	"api-gateway/gen/gpt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GptClient struct {
	conn    *grpc.ClientConn
	service gpt.RecommendationClient
}

func NewGptClient(address string) (*GptClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	service := gpt.NewRecommendationClient(conn)
	return &GptClient{
		conn:    conn,
		service: service,
	}, nil
}

func (c * GptClient) GetGPTRecommendation(ctx context.Context, req *gpt.UserRequest) (*gpt.GPTResponse, error) {
	resp, err := c.service.GetGPTRecommendation(ctx, req)
	if err != nil {
		log.Printf("Failed to get recommendation: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *GptClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Printf("Failed to close gRPC connection: %v", err)
	}
}

