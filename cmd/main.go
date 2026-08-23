package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	// Camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada usecase
	productUseCase := usecase.NewProductUseCase(*ProductRepository)

	// Passando os dados mockados - Camada de controllers
	productController := controller.NewProductController(*productUseCase)

	// Criando uma rota para testar a conectividade
	// o gin.H permite criar um map para envio de mensagem no formato JSON
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Criando a rota para buscar os produtos com dados mockados
	server.GET("/products", productController.GetProducts)

	server.POST("/products", productController.CreateProduct)

	server.GET("/products/:productId", productController.GetProductByID)

	server.DELETE("/products/:productId", productController.DeleteById)

	server.PUT("/products/:productId", productController.UpdateById)

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
