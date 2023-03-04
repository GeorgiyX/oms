package service

import (
	"route256/loms/internal/usecase"
	desc "route256/loms/pkg/loms"
)

type Service struct {
	desc.UnimplementedLomsServer
	useCase usecase.UseCase
}

func New(businessLogic usecase.UseCase) *Service {
	return &Service{
		useCase: businessLogic,
	}
}
