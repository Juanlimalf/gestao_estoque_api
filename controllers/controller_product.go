package controllers

import (
	"database/sql"
	"fmt"
	"main/model"
	"main/repository"
	"main/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	use_case usecase.ProductUseCase
}

func CreateProductController(db *sql.DB) productController {
	productRepository := repository.CreateProductRepository(db)
	productUseCase := usecase.CreateProductUseCase(productRepository)

	return productController{
		use_case: productUseCase,
	}
}

// @Summary Get all products
// @Description Get a list of all products
// @Schemes
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {string} GetProductController
// @Router /products/ [get]
func (p *productController) GetProductController(c *gin.Context) {
	products := p.use_case.GetAllProducts()

	c.JSON(http.StatusOK, products)
}

// @Summary Create a new product
// @Description Create a new product with the input payload
// @Schemes
// @Tags Products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product to create"
// @Success 201 {string} CreateProductController
// @Router /products/ [post]
func (p *productController) CreateProductController(c *gin.Context) {
	var newProduct model.Product

	error := c.BindJSON(&newProduct)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	fmt.Println(newProduct)

	// if err := p.use_case.CreateProduct(newProduct); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusCreated, newProduct)
}
