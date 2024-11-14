package api

import "github.com/front-go/gateway/internal/model"

type SrvI interface {
	Summator(in *model.In)
}
