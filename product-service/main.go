package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/Jvls1/go-ecommerce/product-service/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Product Service
// @version         1.0
// @host      			:8080
// @BasePath  			/api/v1
func main() {
	connectDatabase()
	createWebService()
}

func connectDatabase() {
	cfg := getEnvVariable("DB_CONNECTION")

	var err error
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}
	pingDatabase(db)
	fmt.Println("Connected!")
}

func pingDatabase(db *sql.DB) {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func createWebService() {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine = defineRoutes(engine)
	engine.Run(getEnvVariable("PORT"))
}

func defineRoutes(engine *gin.Engine) *gin.Engine {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/")
	return engine
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
