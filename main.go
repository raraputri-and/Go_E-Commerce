package main

import (
	"e-commerce/customer"
	"e-commerce/handler"
	"e-commerce/initializers"
	"e-commerce/middleware"
	"e-commerce/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	initializers.LoadEnvVariables()
	db, err = initializers.ConnectToDatabase()
	err = initializers.SnycDatabase(db)

	if err != nil {
		log.Fatal("db connection failed")
	}

}

func main() {

	customerRepository := customer.NewRepository(db)
	customerService := customer.NewService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	routerV1 := router.Group("/v1")

	routerV1Customer := routerV1.Group("/customer", middleware.RequireAuth)

	// CUSTOMER
	routerV1Customer.GET("/:id", customerHandler.GetCustomer)
	routerV1Customer.POST("", customerHandler.PostCustomerHandler)
	routerV1Customer.PUT("/:id", customerHandler.UpdateCustomerHandler)

	// USER LOGIN
	routerV1.POST("/signup", userHandler.Signup)
	routerV1.POST("/login", userHandler.Login)

	router.Run(":8081")
}
