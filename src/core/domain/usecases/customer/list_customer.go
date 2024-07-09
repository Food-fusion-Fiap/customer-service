package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ListCustomerUsecase struct {
	CustomerRepository gateways.CustomerRepository
}

func (r *ListCustomerUsecase) Execute(inputDto dtos.ListCustomerDto) (*entities.Customer, error) {
	if inputDto.CPF == "" {
		return nil, nil
	}

	customer := entities.Customer{
		CPF: inputDto.CPF,
	}

	return r.CustomerRepository.FindFirstByCpf(&customer)
}
