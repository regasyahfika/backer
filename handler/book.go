package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"backer/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func response(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          int(b.ID),
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Discount:    b.Discount,
		Rating:      b.Rating,
	}
}

func (h bookHandler) GetBooks(c *gin.Context) {

	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := response(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    booksResponse,
		"message": "Ok",
	})
}

func (h bookHandler) GetDetailBook(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	book, err := h.bookService.FindByID(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// bookResponse := response(book)

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"data":    book,
	})
}

func (h bookHandler) CreateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    book,
		"message": "Status Success",
	})
}

func (h bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    response(book),
		"message": "Status Success",
	})
}

func (h bookHandler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	book, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    response(book),
		"message": "Delete Success",
	})
}
