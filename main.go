package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/initializers"
	"pustaka-api/middleware"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	initializers.LoadEnvVariables()
	db, err = initializers.ConnectToDatabase()
	err = initializers.SyncDatabase(db)

	if err != nil {
		log.Fatal("Connection to database failed")
	}
}

func main() {

	// // ADD NEW BOOK HARDCODE
	// book := book.Book{}
	// book.Title = "Atomic habits"
	// book.Price = 12000
	// book.Description = "Buku self development tentang membangun kebiasaan baik"
	// book.Rating = 4

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book record")
	// }

	// // UPDATE BOOK
	// var book book.Book

	// err = db.Debug().First(&book).Error
	// // err = db.Debug().Last(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// } else {
	// 	book.Title = "Atomic habitsa"
	// 	err = db.Save(&book).Error
	// 	if err != nil {
	// 		fmt.Println("Error updating book record")
	// 	}
	// }

	// // SEARCH BY QUERY
	// var book book.Book
	// err = db.Debug().First(&book, 2).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }

	// fmt.Println("book object %v", book)

	// DELETE BY HARDCODE ID
	// var book book.Book
	// book.ID = 1
	// err = db.Delete(&book).Error

	router := gin.Default()

	routerV1 := router.Group("/v1")
	routerV1Books := routerV1.Group("/books", middleware.RequireAuth)

	// no func implemented
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name": "Fairuz Satria",
	// 		"bio":  "A Software Engineer & content creator",
	// 	})
	// })

	// router.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"title": "Hello World",
	// 	})
	// })

	// ROUTE GROUP
	// routerV1.GET("/", handler.RootHandler)

	// routerV1.GET("/hello", handler.HelloHandler)

	// routerV1.GET("books/:id/:title", handler.BooksHandler)

	// routerV1.GET("/query", handler.QueryHandler)

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// ROUTE Login
	routerV1.POST("/signup", userHandler.Signup)
	routerV1.POST("/login", userHandler.Login)
	// ROUTE BOOKS
	routerV1Books.POST("", bookHandler.PostBooksHandler)
	routerV1Books.GET("", bookHandler.GetBooks)
	routerV1Books.GET("/:id", bookHandler.GetBook)
	routerV1Books.PUT("/:id", bookHandler.UpdateBookHandler)
	routerV1Books.DELETE("/:id", bookHandler.DeleteBook)
	routerV1Books.DELETE("/:id", bookHandler.DeleteBook)

	// FOR RUNNING
	router.Run(":9000")

}
