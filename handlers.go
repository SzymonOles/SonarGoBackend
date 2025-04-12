package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Produkty
func getProducts(c echo.Context) error {
	var products []Product
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func getProductsByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}
	var products []Product
	db.Scopes(ScopeByCategory(uint(categoryID))).Find(&products)
	return c.JSON(http.StatusOK, products)
}

func getExpensiveProducts(c echo.Context) error {
	var products []Product
	db.Scopes(ScopeExpensiveProducts(100.0)).Find(&products)
	return c.JSON(http.StatusOK, products)
}

func createProduct(c echo.Context) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	db.Create(product)
	return c.JSON(http.StatusCreated, product)
}

func updateProduct(c echo.Context) error {
	id := c.Param("id")
	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	if err := c.Bind(&product); err != nil {
		return err
	}
	db.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id := c.Param("id")
	db.Delete(&Product{}, id)
	return c.NoContent(http.StatusNoContent)
}

// Kategorie
func getCategories(c echo.Context) error {
	var categories []Category
	db.Preload("Products").Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func createCategory(c echo.Context) error {
	category := new(Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	db.Create(category)
	return c.JSON(http.StatusCreated, category)
}

func updateCategory(c echo.Context) error {
	id := c.Param("id")
	var category Category
	db.First(&category, id)
	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, "Category not found")
	}
	if err := c.Bind(&category); err != nil {
		return err
	}
	db.Save(&category)
	return c.JSON(http.StatusOK, category)
}

func deleteCategory(c echo.Context) error {
	id := c.Param("id")
	db.Delete(&Category{}, id)
	return c.NoContent(http.StatusNoContent)
}

// Koszyk
func getCarts(c echo.Context) error {
	var carts []Cart
	db.Preload("Products").Find(&carts)
	return c.JSON(http.StatusOK, carts)
}

func createCart(c echo.Context) error {
	cart := new(Cart)
	if err := c.Bind(cart); err != nil {
		return err
	}
	db.Create(cart)
	return c.JSON(http.StatusCreated, cart)
}

func deleteCart(c echo.Context) error {
	id := c.Param("id")
	var cart Cart
	db.First(&cart, id)
	if cart.ID == 0 {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	db.Delete(&cart)
	return c.NoContent(http.StatusNoContent)
}

func addProductToCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))
	productID, _ := strconv.Atoi(c.Param("product_id"))

	var cart Cart
	db.First(&cart, cartID)
	if cart.ID == 0 {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	var product Product
	db.First(&product, productID)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	db.Model(&cart).Association("Products").Append(&product)

	return c.JSON(http.StatusOK, cart)
}

func removeProductFromCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))
	productID, _ := strconv.Atoi(c.Param("product_id"))

	var cart Cart
	db.First(&cart, cartID)
	if cart.ID == 0 {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	var product Product
	db.First(&product, productID)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	db.Model(&cart).Association("Products").Delete(&product)

	return c.JSON(http.StatusOK, cart)
}
