package controllers

import (
	"database/sql"
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

func (p *productController) GetProductController(c *gin.Context) {
	products := p.use_case.GetAllProducts()

	c.JSON(http.StatusOK, products)
}
