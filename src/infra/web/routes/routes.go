package routes

import (
	"log"

	controllers "github.com/CAVAh/api-tech-challenge/src/adapters/controllers/customer"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/customer"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/database"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()
	customerRepository := &repositories.CustomerRepository{
		DB: database.DB,
	}

	lgpdRemovalRequestRepository := &repositories.LgpdRemovalRequestRepository{
		DB: database.DB,
	}

	listUsecase := &usecases.ListCustomerUsecase{CustomerRepository: customerRepository}
	createUsecase := &usecases.CreateCustomerUsecase{CustomerRepository: customerRepository}
	createLgpdRemovalRequest := &usecases.CreateLgpdRemovalRequestUsecase{LgpdRemovalRequestRepository: lgpdRemovalRequestRepository}

	router.GET("/customers", func(c *gin.Context) {
		controllers.ListCustomers(c, listUsecase)
	})

	router.POST("/customers", func(c *gin.Context) {
		controllers.CreateCustomer(c, createUsecase)
	})

	router.POST("/customers/lgpd-exclusion", func(c *gin.Context) {
		controllers.CreateLgpdRemovalRequestCustomer(c, createLgpdRemovalRequest)
	})

	err := router.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
