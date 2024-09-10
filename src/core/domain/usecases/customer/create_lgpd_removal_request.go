package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CreateLgpdRemovalRequestUsecase struct {
	LgpdRemovalRequestRepository gateways.LgpdRemovalRequestRepository
}

func (r *CreateLgpdRemovalRequestUsecase) Execute(inputDto dtos.LGPDRemovalRequestDto) (*entities.LgpdRemovalRequest, error) {
	removalRequest := entities.LgpdRemovalRequest{
		Name:           inputDto.Name,
		Address:        inputDto.Address,
		PhoneNumber:    inputDto.PhoneNumber,
		PaymentHistory: inputDto.PaymentHistory,
	}

	return r.LgpdRemovalRequestRepository.Create(&removalRequest)
}
