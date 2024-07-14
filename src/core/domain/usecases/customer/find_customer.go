package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/utils"
)

type ListCustomerUsecase struct {
	CustomerRepository gateways.CustomerRepository
}

func (r *ListCustomerUsecase) Execute(inputDto dtos.ListCustomerDto) (string, error) {
	var customer entities.Customer

	// Se o CPF for vazio, gere o token com customerId nulo
	if inputDto.CPF == "" {
		token, err := utils.GenerateJWT(nil)

		if err != nil {
			return "", err
		}

		return token, nil
	}

	customer.CPF = inputDto.CPF

	// Se o cliente existir, gere o token com o customerId do cliente
	foundCustomer, err := r.CustomerRepository.FindFirstByCpf(&customer)

	if err == nil {
		token, err := utils.GenerateJWT(foundCustomer.ID)

		if err != nil {
			return "", err
		}

		return token, nil
	}

	// Se o cliente não existir, gere o token com customerId nulo
	token, err := utils.GenerateJWT(nil)

	if err != nil {
		return "", err
	}

	return token, nil
}
