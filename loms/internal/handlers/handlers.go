package handlers

import (
	"route256/loms/internal/usecase"
)

type Handler struct {
	useCase usecase.UseCase
}

func New(businessLogic usecase.UseCase) *Handler {
	return &Handler{
		useCase: businessLogic,
	}
}
