package main

import (
	"e-commerce/customer"
	"e-commerce/handler"
	"e-commerce/initializers"
	"e-commerce/middleware"
	"e-commerce/product"
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

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	routerV1 := router.Group("/v1")

	// CUSTOMER
	routerV1Customer := routerV1.Group("/customer", middleware.RequireAuth)

	routerV1Customer.GET("", customerHandler.GetCustomers)
	routerV1Customer.GET("/:id", customerHandler.GetCustomer)
	routerV1Customer.POST("", customerHandler.PostCustomerHandler)
	routerV1Customer.PUT("/:id", customerHandler.UpdateCustomerHandler)

	// PRODUCT
	routerV1Product := routerV1.Group("/product", middleware.RequireAuth)
	routerV1Product.GET("", productHandler.GetProducts)
	routerV1Product.GET("/:id", productHandler.GetProduct)
	routerV1Product.POST("", productHandler.PostProductHandler)
	routerV1Product.PUT("/:id", productHandler.UpdateProductHandler)
	routerV1Product.DELETE("/:id", productHandler.DeleteProduct)

	// USER LOGIN
	routerV1.POST("/signup", userHandler.Signup)
	routerV1.POST("/login", userHandler.Login)

	router.Run(":8081")
}
