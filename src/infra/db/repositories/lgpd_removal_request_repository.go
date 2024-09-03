package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/database"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/models"
	"strings"
)

type LgpdRemovalRequestRepository struct {
	DB database.Database
}

func (r LgpdRemovalRequestRepository) Create(removalRequest *entities.LgpdRemovalRequest) (*entities.LgpdRemovalRequest, error) {

	removalRequestModel := models.LgpdRemovalRequest{
		Name:           removalRequest.Name,
		Address:        removalRequest.Address,
		PhoneNumber:    removalRequest.PhoneNumber,
		PaymentHistory: removalRequest.PaymentHistory,
	}

	if err := r.DB.Create(&removalRequestModel); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("cliente já existe no sistema")
		} else {
			return nil, errors.New("ocorreu um erro desconhecido ao criar o cliente")
		}
	}

	result := removalRequestModel.ToDomain()

	return &result, nil
}
