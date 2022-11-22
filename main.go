package main

import (
	"net/http"

	"github.com/Taras-Rm/aws_rds/repository"
	"github.com/Taras-Rm/aws_rds/setup"
	"github.com/gin-gonic/gin"
)

func main() {
	db := setup.GetPostgresConnect()

	productRepo := repository.NewProductRepository(db)

	routes := gin.Default()

	routes.POST("api/products", createProduct(productRepo))
	routes.GET("api/products", getAllProducts(productRepo))

	routes.Run(":8080")

}

type createProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func createProduct(repo repository.ProductInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inp createProductInput
		err := ctx.BindJSON(&inp)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		product := repository.ProductModel{
			Name:  inp.Name,
			Price: uint64(inp.Price),
		}

		err = repo.Create(product)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "created"})
	}
}

func getAllProducts(repo repository.ProductInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := repo.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success", "products": products})
	}
}
