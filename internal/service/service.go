package service

import "github.com/front-go/gateway/internal/model"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Summator(in *model.In) {
	in.Sum = in.A + in.B
}
