package api

import (
	"context"
	"github.com/front-go/gateway/internal/model"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Passport struct {
	Number string `json:"number"`
	Series string `json:"series"`
}

type FullName struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type AuthClient interface {
	DoSignup(ctx context.Context, signupInfo model.UserSignup) (bool, error)
}
