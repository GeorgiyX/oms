package handlers

import (
	"route256/checkout/internal/usecase"
)

type Handler struct {
	businessLogic usecase.UseCase
}

func New(businessLogic usecase.UseCase) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}
