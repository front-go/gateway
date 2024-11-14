package api

import (
	"file-structure/internal/model"
)

type Handler struct {
	srv SrvI
}

func NewHandler(srv SrvI) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) Handle(x, y int) int {
	inModel := &model.In{
		A: x,
		B: y,
	}
	h.srv.Summator(inModel)
	return inModel.Sum
}
