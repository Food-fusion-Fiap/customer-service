package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type LgpdRemovalRequestRepository interface {
	Create(removalRequest *entities.LgpdRemovalRequest) (*entities.LgpdRemovalRequest, error)
}
