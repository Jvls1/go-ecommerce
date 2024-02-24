package main

import (
	"database/sql"
	_ "github.com/Jvls1/go-ecommerce/product-service/docs"
	"github.com/Jvls1/go-ecommerce/product-service/internal/repository"
	"github.com/Jvls1/go-ecommerce/product-service/internal/service"
	"github.com/Jvls1/go-ecommerce/product-service/internal/webserver"
	"github.com/Jvls1/go-ecommerce/product-service/pkg"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Product Service
// @version         1.0
// @host      		:8080
// @BasePath  		/api/v1
func main() {
	loadEnvVariables()
	db := pkg.NewDBConnection()
	productRepo := createRepoInstance(db)
	productService := createServiceInstance(productRepo)
	productHandler := createHandlerInstance(productService)
	createWebService(productHandler)
}

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createRepoInstance(db *sql.DB) repository.ProductRepository {
	productRepo := repository.NewProductRepository(db)
	return productRepo
}

func createServiceInstance(productRepo repository.ProductRepository) service.ProductService {
	productService := service.NewProductService(productRepo)
	return productService
}

func createHandlerInstance(productService service.ProductService) webserver.ProductHandler {
	productHandler := webserver.NewProductHandler(productService)
	return productHandler
}

func createWebService(productHandler webserver.ProductHandler) {
	engine := gin.Default()
	err := engine.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Error trying to run the project: %s", err)
	}
	engine = defineRoutes(engine, productHandler)
	err = engine.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error trying to run the project: %s", err)
	}
}

func defineRoutes(engine *gin.Engine, productHandler webserver.ProductHandler) *gin.Engine {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := engine.Group("/api/v1/products")
	{
		v1.POST("/", productHandler.CreateProduct)
		v1.GET("/", productHandler.GetProducts)
		v1.GET("/:id", productHandler.GetProductById)
		v1.GET("/department/:departmentId", productHandler.GetProductByDepartmentID)
	}
	return engine
}
