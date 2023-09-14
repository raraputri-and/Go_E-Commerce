package handler

import (
	"e-commerce/product"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) PostProductHandler(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s",
					e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
	}

	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)
	product, err := h.productService.Create(productRequest, uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func (h *productHandler) GetProducts(c *gin.Context) {
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	customerID, _ := claims["sub"].(float64)

	products, err := h.productService.FindAll(uint(customerID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var productResponses []product.ProductResponse

	for _, b := range products {

		// COBA NEW CODE FOR THIS METHOD
		productResponses = append(productResponses, product.ConvertToProductResponse(b))

		// productResponse := product.ConvertToProductResponse(b)
		// productResponses = append(productResponses, productResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func (h *productHandler) GetProduct(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.productService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productResponse := product.ConvertToProductResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (h *productHandler) UpdateProductHandler(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBindJSON(&productRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s",
					e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}
	}

	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.productService.Update(ID, productRequest)
	productResponse := product.ConvertToProductResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.productService.Delete(ID)
	productResponse := product.ConvertToProductResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (h *productHandler) SearchByName(c *gin.Context) {
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	customerID, _ := claims["sub"].(float64)
	name := c.Query("name")

	products, err := h.productService.SearchByName(uint(customerID),string(name))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var productResponses []product.ProductResponse

	for _, b := range products {

		// COBA NEW CODE FOR THIS METHOD
		productResponses = append(productResponses, product.ConvertToProductResponse(b))

		// productResponse := product.ConvertToProductResponse(b)
		// productResponses = append(productResponses, productResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}