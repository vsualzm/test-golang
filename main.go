package main

import (
	"fmt"
	"test-golang/config"
	"test-golang/handler"
	"test-golang/repository"
	"test-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// start API

	fmt.Println("Started API")

	config.ConnectionDatabase()

	// Setup repository, service, and handler
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.GET("/products", productHandler.GetAllProducts)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	// cartService := service.CartService{DB: d}
	// cartHandler := handler.CartHandler{CartService: &cartService}

	// // router := gin.Default()
	// router.POST("/cart/:cart_id/add", cartHandler.AddToCart)
	// router.GET("/cart/:cart_id", cartHandler.ViewCart)
	// router.POST("/cart/:cart_id/checkout", cartHandler.Checkout)

	router.Run(":8080")
}
