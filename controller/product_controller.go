package controller

import (
	"go-api/usecase"
	"net/http"

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
