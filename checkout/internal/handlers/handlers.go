package handlers

import (
	"route256/checkout/internal/usecase"
)

type Handler struct {
	useCase usecase.UseCase
}

func New(businessLogic usecase.UseCase) *Handler {
	return &Handler{
		useCase: businessLogic,
	}
}
