package checkout

import (
	"route256/checkout/internal/usecase"
	desc "route256/checkout/pkg/checkout"
)

type Service struct {
	desc.UnimplementedCheckoutServer
	useCase usecase.UseCase
}

func New(businessLogic usecase.UseCase) *Service {
	return &Service{
		useCase: businessLogic,
	}
}
