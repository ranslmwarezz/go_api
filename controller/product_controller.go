package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	// UseCase
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) *productController {
	return &productController{productUseCase: usecase}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	// Exemplo de como mockar dados

	/* productsMock := []model.Product{
		{
			ID:    1,
			Name:  "teste",
			Price: 35.00,
		},
	}
	*/

	products, err := p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(200, products)

}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		fmt.Println("Error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)

}

func (p *productController) GetProductByID(ctx *gin.Context) {

	id := ctx.Param("productId")

	var response model.Response

	if id == "" {
		response.Message = "Product ID cannot be null"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = "The product ID must be a number"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductByID(productId)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if product == nil {
		response.Message = "Product not found"
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (p *productController) DeleteById(ctx *gin.Context) {

	id := ctx.Param("productId")

	var response model.Response

	if id == "" {
		response.Message = "Product ID cannot be null"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = "The product ID must be a number"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product := p.productUseCase.DeleteById(productId)

	if product == nil {
		response.Message = "Product not found"
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (p *productController) UpdateById(ctx *gin.Context) {

	id := ctx.Param("productId")

	var response model.Response

	if id == "" {
		response.Message = "Product ID cannot be null"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = "The product ID must be a number"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product

	err = ctx.BindJSON(&product)
	if err != nil {
		response.Message = "Invalid product data"
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.UpdateById(productId, product)
	if err != nil {
		response.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "Product updated successfully"
	ctx.JSON(http.StatusOK, response)

}
