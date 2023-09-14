package handler

import (
	"e-commerce/customer"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type customerHandler struct {
	customerService customer.Service
}

func NewCustomerHandler(service customer.Service) *customerHandler {
	return &customerHandler{service}
}

func (h *customerHandler) GetCustomers(c *gin.Context) {
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)
	
	customers, err := h.customerService.FindAll(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var customersResponse []customer.CustomerResponse

	for _, c := range customers {
		customerResponse := customer.ConvertToCustomerResponse(c)
		customersResponse = append(customersResponse, customerResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customersResponse,
	})
}

func (h *customerHandler) GetCustomer(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	b,err := h.customerService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}

	customerResponse := customer.ConvertToCustomerResponse(b)

	c.JSON(http.StatusBadRequest, gin.H{
		"data": customerResponse,
	})
}

func (h *customerHandler) PostCustomerHandler(c *gin.Context) {
	var customerRequest customer.CustomerRequest
	err := c.ShouldBindJSON(&customerRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)

	customer, err := h.customerService.Create(customerRequest, uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})

}

func (h *customerHandler) UpdateCustomerHandler(c *gin.Context){
	var customerRequest customer.CustomerRequest

	err := c.ShouldBindJSON(&customerRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _,e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors" : errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.customerService.Update(ID, customerRequest)
	customerResponse := customer.ConvertToCustomerResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customerResponse,
	})
}