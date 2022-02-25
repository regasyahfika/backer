package main

import (
	"backer/book"
	"backer/database"
	"backer/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := database.Connect(&gorm.DB{})
	database.Migrate(db)

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewHandler(bookService)

	router := gin.Default()

	v1 := router.Group("v1")
	v1.GET("/book", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetDetailBook)
	v1.POST("/book", bookHandler.CreateBookHandler)
	v1.PUT("/book/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/book/:id", bookHandler.DeleteBook)
	router.Run()

}
