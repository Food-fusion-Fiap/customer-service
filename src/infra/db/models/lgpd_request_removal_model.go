package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/utils"
	"gorm.io/gorm"
)

type LgpdRemovalRequest struct {
	gorm.Model
	Name           string
	Address        string
	PhoneNumber    string
	PaymentHistory bool
}

func (c LgpdRemovalRequest) ToDomain() entities.LgpdRemovalRequest {
	return entities.LgpdRemovalRequest{
		ID:             c.ID,
		Name:           c.Name,
		Address:        c.Address,
		PhoneNumber:    c.PhoneNumber,
		PaymentHistory: c.PaymentHistory,
		CreatedAt:      c.CreatedAt.Format(utils.CompleteEnglishDateFormat),
	}
}
