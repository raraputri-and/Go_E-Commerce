package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		// c.JSON(http.StatusBadRequest, err)
		// fmt.Println(err)
		// return
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
	book, err := h.bookService.Create(bookRequest, uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"title":     bookInput.Title,
	// 	"price":     bookInput.Price,
	// 	"sub_title": bookInput.SubTiltle,
	// })
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)

	books, err := h.bookService.FindAll(uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.ConverToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.bookService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := book.ConverToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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
	b, err := h.bookService.Update(ID, bookRequest)
	bookResponse := book.ConverToBookResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.bookService.Delete(ID)
	bookResponse := book.ConverToBookResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}
