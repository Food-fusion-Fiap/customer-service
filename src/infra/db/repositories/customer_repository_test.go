package repositories

import (
	"errors"
	"testing"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/mocks"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)
	repo := CustomerRepository{DB: mockDB}

	entity := &entities.Customer{
		Name:  "John Doe",
		CPF:   "12345678901",
		Email: "john@example.com",
	}

	mockDB.EXPECT().Create(gomock.Any()).Return(nil)

	result, err := repo.Create(entity)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "John Doe", result.Name)
}

func TestCreateCustomer_Duplicate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)
	repo := CustomerRepository{DB: mockDB}

	entity := &entities.Customer{
		Name:  "John Doe",
		CPF:   "12345678901",
		Email: "john@example.com",
	}

	mockDB.EXPECT().Create(gomock.Any()).Return(errors.New("duplicate key value violates unique constraint"))

	result, err := repo.Create(entity)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "cliente já existe no sistema", err.Error())
}

func TestCreateCustomer_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)
	repo := CustomerRepository{DB: mockDB}

	entity := &entities.Customer{
		Name:  "John Doe",
		CPF:   "12345678901",
		Email: "john@example.com",
	}

	mockDB.EXPECT().Create(gomock.Any()).Return(errors.New("some error"))

	result, err := repo.Create(entity)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "ocorreu um erro desconhecido ao criar o cliente", err.Error())
}

func TestListCustomers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)
	repo := CustomerRepository{DB: mockDB}

	entity := &entities.Customer{
		Name:  "John Doe",
		CPF:   "12345678901",
		Email: "john@example.com",
	}

	mockDB.EXPECT().Find(gomock.Any()).Return(nil)

	result, err := repo.List(entity)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
