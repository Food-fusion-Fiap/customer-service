package controllers

import (
	"fmt"
	"net/http"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/customer"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func ListCustomers(c *gin.Context, usecase *usecases.ListCustomerUsecase) {
	var inputDto dtos.ListCustomerDto

	queryParams := c.Request.URL.Query().Encode()
	fmt.Printf("Query string recebida: %s\n", queryParams)

	c.BindQuery(&inputDto)

	fmt.Printf("inputDto após o bind: %+v\n", inputDto)

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	result, err := usecase.Execute(inputDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateCustomer(c *gin.Context, usecase *usecases.CreateCustomerUsecase) {
	var inputDto dtos.CreateCustomerDto

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	result, err := usecase.Execute(inputDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
