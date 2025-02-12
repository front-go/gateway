package auth

import (
	"context"
	"github.com/front-go/auth/pkg/auth"
	"github.com/front-go/gateway/internal/model"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	client auth.AuthServiceClient
}

func NewClient() *Client {
	conn, err := grpc.NewClient("localhost:8095", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := auth.NewAuthServiceClient(conn)
	return &Client{
		client: client,
	}
}

func (c *Client) DoSignup(ctx context.Context, signupInfo model.UserSignup) (bool, error) {
	res, err := c.client.Signup(ctx, &auth.SignupIn{
		Username:        signupInfo.Login,
		Password:        signupInfo.Password,
		ConfirmPassword: signupInfo.ConfirmPassword,
	})
	if err != nil {
		return false, err
	}
	return res.Success, nil
}
