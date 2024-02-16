package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/Jvls1/go-ecommerce/product-service/docs"
	"github.com/Jvls1/go-ecommerce/product-service/internal/repository"
	"github.com/Jvls1/go-ecommerce/product-service/internal/service"
	"github.com/Jvls1/go-ecommerce/product-service/internal/webserver"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Product Service
// @version         1.0
// @host      		:8080
// @BasePath  		/api/v1
func main() {
	db := connectDatabase()
	productRepo := createRepoInstace(db)
	productService := createServiceInstace(productRepo)
	productHandler := createHandlerInstace(productService)
	createWebService(productHandler)
}

func connectDatabase() *sql.DB {
	dbName := getEnvVariable("DB_NAME")
	dbUser := getEnvVariable("DB_USER")
	dbPassword := getEnvVariable("DB_PASSWORD")
	dbHost := getEnvVariable("DB_HOST")
	dbPort := getEnvVariable("DB_PORT")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	pingDatabase(db)
	return db
}

func pingDatabase(db *sql.DB) {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func createRepoInstace(db *sql.DB) *repository.ProductRepo {
	productRepo := repository.NewProductRepo(db)
	return productRepo
}

func createServiceInstace(productRepo *repository.ProductRepo) *service.ProductService {
	productService := service.NewProductService(productRepo)
	return productService
}

func createHandlerInstace(productService *service.ProductService) *webserver.ProductHandler {
	productHandler := webserver.NewProductHandler(productService)
	return productHandler
}

func createWebService(productHandler *webserver.ProductHandler) {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine = defineRoutes(engine, productHandler)
	engine.Run(getEnvVariable("PORT"))
}

func defineRoutes(engine *gin.Engine, productHandler *webserver.ProductHandler) *gin.Engine {
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

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
