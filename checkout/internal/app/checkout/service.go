package checkout

import (
	"route256/checkout/internal/usecases/checkout"
	desc "route256/checkout/pkg/checkout"
)

type Service struct {
	desc.UnimplementedCheckoutServer
	useCase checkout.UseCase
}

func New(businessLogic checkout.UseCase) *Service {
	return &Service{
		useCase: businessLogic,
	}
}
