package service

import "file-structure/internal/model"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Summator(in *model.In) {
	in.Sum = in.A + in.B
}
