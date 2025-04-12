package main

import "github.com/labstack/echo/v4"

func registerRoutes(e *echo.Echo) {
	// Produkty
	e.GET("/products", getProducts)
	e.GET("/products/category/:category_id", getProductsByCategory)
	e.GET("/products/expensive", getExpensiveProducts)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	// Kategorie
	e.GET("/categories", getCategories)
	e.POST("/categories", createCategory)
	e.PUT("/categories/:id", updateCategory)
	e.DELETE("/categories/:id", deleteCategory)

	// Koszyk
	e.GET("/carts", getCarts)
	e.POST("/carts", createCart)
	e.DELETE("/carts/:id", deleteCart)

	// Produkty w Koszyku
	e.POST("/carts/:cart_id/products/:product_id", addProductToCart)
	e.DELETE("/carts/:cart_id/products/:product_id", removeProductFromCart)
}
