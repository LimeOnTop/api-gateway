package client

import (
	"context"
	"log"

	"api-gateway/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	conn    *grpc.ClientConn
	service auth.AuthClient
}

func NewAuthClient(address string) (*AuthClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	service := auth.NewAuthClient(conn)
	return &AuthClient{
		conn:    conn,
		service: service,
	}, nil
}

func (c *AuthClient) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	resp, err := c.service.Register(ctx, req)
	if err != nil {
		log.Printf("Failed to register: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *AuthClient) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	resp, err := c.service.Login(ctx, req)
	if err != nil {
		log.Printf("Failed to login: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *AuthClient) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	resp, err := c.service.RefreshToken(ctx, req)
	if err != nil {
		log.Printf("Failed to refresh token: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *AuthClient) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {

	resp, err := c.service.ValidateToken(ctx, req)
	if err != nil {
		log.Printf("Failed to validate token: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *AuthClient)  Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	resp, err := c.service.Logout(ctx, req)
	if err != nil {
		log.Printf("Failed to logout: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *AuthClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Printf("Failed to close gRPC connection: %v", err)
	}
}
