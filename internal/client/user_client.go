package client

import (
	"context"
	"log"

	"api-gateway/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	conn *grpc.ClientConn
	service user.UserServiceClient
}

func NewUserClient(address string) (*UserClient, error) {
	conn, err:= grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	service := user.NewUserServiceClient(conn)
	return &UserClient{
		conn: conn,
		service: service,
	}, nil
}

func (c* UserClient) GetUserProducts(ctx context.Context, req *user.UserRequest) (*user.GetProductsResponse, error) {
	resp, err := c.service.GetUserProducts(ctx, req)
	if err != nil {
		log.Printf("Failed to get user products: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *UserClient) GetUserPreference(ctx context.Context, req *user.UserRequest) (*user.GetPreferenceResponse, error) {
	resp, err := c.service.GetUserPreference(ctx, req)
	if err != nil {
		log.Printf("Failed to get user preference: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *UserClient) AddUserProduct(ctx context.Context, req *user.AddProductRequest) (*user.AddProductResponse, error) {
	resp, err := c.service.AddUserProduct(ctx, req)
	if err != nil {
		log.Printf("Failed to add user product: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *UserClient) RemoveUserProduct(ctx context.Context, req *user.RemoveProductRequest) (*user.RemoveProductResponse, error) {
	resp, err := c.service.RemoveUserProduct(ctx, req)
	if err != nil {
		log.Printf("Failed to remove user product: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *UserClient) UpdateUserPreference(ctx context.Context, req *user.UpdatePreferenceRequest) (*user.UpdatePreferenceResponse, error) {
	resp, err := c.service.UpdateUserPreference(ctx, req)
	if err != nil {
		log.Printf("Failed to update user preference: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *UserClient) RemoveUserPreference(ctx context.Context, req *user.RemovePreferenceRequest) (*user.RemovePreferenceResponse, error) {
	resp, err := c.service.RemoveUserPreference(ctx, req)
	if err != nil {
		log.Printf("Failed to remove user preference: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *UserClient) Close() {
	if err:= c.conn.Close(); err != nil {
		log.Printf("Failed to close connection: %v", err)
	}
}
