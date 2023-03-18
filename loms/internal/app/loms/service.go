package loms

import (
	"route256/loms/internal/usecase/loms"
	desc "route256/loms/pkg/loms"
)

type Service struct {
	desc.UnimplementedLomsServer
	useCase loms.UseCase
}

func New(businessLogic loms.UseCase) *Service {
	return &Service{
		useCase: businessLogic,
	}
}
